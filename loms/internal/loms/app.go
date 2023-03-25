package loms

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/pkg/errors"
	desc "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/loms_v1"
	lomsV1 "gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/api/loms_v1"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/config"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/interceptors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	impl            *lomsV1.Implementation
	serviceProvider *serviceProvider

	config *config.ConfigData

	grpcServer *grpc.Server
}

/// INITIALIZATION ///

func New(ctx context.Context, data *config.ConfigData) (*App, error) {
	app := App{
		config: data,
	}

	err := app.Init(ctx)

	return &app, err
}

func (app *App) Init(ctx context.Context) error {
	inits := []func(context.Context) error{
		app.initServiceProvider,
		app.initServer,
		app.initGRPCServer,
		//app.initHTTPHandlers,
	}

	for _, fn := range inits {
		if err := fn(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (app *App) initServiceProvider(ctx context.Context) error {
	app.serviceProvider = NewServiceProvider(app.config)

	return nil
}

func (app *App) initServer(ctx context.Context) error {
	service := app.serviceProvider.GetService(ctx)

	app.impl = lomsV1.New(service)

	return nil
}

func (app *App) initGRPCServer(ctx context.Context) error {
	app.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(interceptors.LoggingInterceptor),
		),
	)

	reflection.Register(app.grpcServer)

	desc.RegisterLomsV1Server(app.grpcServer, app.impl)

	return nil
}

/// RUN ///

func (app *App) Run() error {
	defer func() {
		app.serviceProvider.db.Close()
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
