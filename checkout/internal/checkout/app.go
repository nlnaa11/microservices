package checkout

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/pkg/errors"
	checkoutV1 "gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/api/checkout_v1"
	grpcLoms "gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/grpc/loms"
	grpcProduct "gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/grpc/product"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/config"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/interceptors"
	desc "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/checkout_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type Client uint8

const (
	CltLoms Client = iota
	CltProduct
)

type App struct {
	impl            *checkoutV1.Implementation
	serviceProvider *serviceProvider

	config *config.ConfigData

	grpcServer *grpc.Server

	clientsConn map[Client]*grpc.ClientConn
	loms        grpcLoms.Client
	product     grpcProduct.Client
}

/// INITIALIZATION ///

func New(ctx context.Context, data *config.ConfigData) (*App, error) {
	app := App{
		config: data,

		clientsConn: make(map[Client]*grpc.ClientConn),
	}

	err := app.Init(ctx)

	return &app, err
}

func (app *App) Init(ctx context.Context) error {
	err := app.initServiceProvider(ctx)
	if err != nil {
		return err
	}

	if err = app.initClient(ctx, CltLoms); err != nil {
		return err
	}

	if err = app.initClient(ctx, CltProduct); err != nil {
		return err
	}

	if err = app.initServer(ctx); err != nil {
		return err
	}

	if err = app.initGRPCServer(ctx); err != nil {
		return err
	}

	//TODO: app.initHTTPHandlers,

	return nil
}

func (app *App) initServiceProvider(_ context.Context) error {
	app.serviceProvider = NewServiceProvider(app.config)

	return nil
}

func (app *App) initClientConn(ctx context.Context, client Client) error {
	var target string

	switch client {
	case CltLoms:
		target = app.config.GetServiceAddress(config.Loms)
	case CltProduct:
		target = app.config.GetServiceAddress(config.Product)
	default:
		target = ""
	}

	if target == "" {
		return errors.New("empty target for client")
	}

	clientConn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return errors.WithMessage(err, "connecting client to server")
	}

	app.clientsConn[client] = clientConn

	return nil
}

func (app *App) initClient(ctx context.Context, client Client) error {
	if err := app.initClientConn(ctx, client); err != nil {
		return err
	}

	switch client {
	case CltLoms:
		app.loms = grpcLoms.New(app.clientsConn[client])
	case CltProduct:
		app.product = grpcProduct.New(app.clientsConn[CltProduct], app.config)
	default:
		return errors.New("unknown client")
	}

	return nil
}

func (app *App) initServer(ctx context.Context) error {
	service := app.serviceProvider.GetService(ctx)

	app.impl = checkoutV1.New(service, app.loms, app.product)

	return nil
}

func (app *App) initGRPCServer(ctx context.Context) error {
	app.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(interceptors.LoggingInterceptor),
		),
	)
	// для постмана: распознаем апи, используемый сервером
	reflection.Register(app.grpcServer)

	// регистрируем поведение
	desc.RegisterCheckoutV1Server(app.grpcServer, app.impl)

	return nil
}

/// RUN ///

func (app *App) Run() error {
	defer func() {
		app.serviceProvider.db.Close()

		for _, client := range app.clientsConn {
			client.Close()
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	if err := app.runGRPC(wg); err != nil {
		return err
	}

	if err := app.runHTTP(wg); err != nil {
		return err
	}

	wg.Wait()

	return nil
}

func (app *App) runGRPC(wg *sync.WaitGroup) error {
	grpcAddr := app.config.GetCommunicationAddress(config.GRPC)
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return errors.WithMessage(err, "failed to create listener")
	}

	go func() {
		defer wg.Done()

		if err = app.grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to process grpc server: %s", err.Error())
		}
	}()

	log.Printf("grpc listening at %v\n", grpcAddr)

	return nil
}

func (app *App) runHTTP(wg *sync.WaitGroup) error {
	httpAddr := app.config.GetCommunicationAddress(config.HTTP)

	go func() {
		defer wg.Done()

		if err := http.ListenAndServe(httpAddr, nil); err != nil {
			log.Fatalf("unabled to listen to http: %s", err.Error())
		}
	}()

	log.Printf("http listening at %v\n", httpAddr)

	return nil
}
