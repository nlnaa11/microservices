package main

import (
	"log"
	"net"
	"net/http"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/api/checkout_v0/addtocart"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/api/checkout_v0/cartlist"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/api/checkout_v0/deletefromcart"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/api/checkout_v0/purchase"
	checkoutV1 "gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/api/checkout_v1"
	grpcLoms "gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/grpc/loms"
	grpcProduct "gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/grpc/product"
	jsonLoms "gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/jsonrpc/loms"
	jsonProduct "gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/jsonrpc/product"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/config"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/interceptors"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/service/checkout"
	desc "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/checkout_v1"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/storage"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/storage/memory"
	"gitlab.ozon.dev/nlnaa/homework-1/libs/wrappers/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	// storage
	mem, err := memory.Init()
	if err != nil {
		log.Fatal("failed to access storage", err)
	}
	stor := storage.New(mem)

	// config
	err = config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	// business logic with behavior
	checkoutService := checkout.New(stor)
	if checkoutService == nil {
		log.Fatal("failed to init checkout service")
	}

	/// grpc communication
	initGrpcCommunication(checkoutService, stor)

	/// json_rpc communication
	initJsonRpcCommunication(checkoutService, stor)
}

func initGrpcCommunication(checkoutService *checkout.Service, stor *storage.WrapStorage) {
	/// 1. Creating connections with clients

	// loms client
	lomsAdd := config.ConfigData.GetServiceAddress(config.Loms)
	if lomsAdd == "" {
		log.Fatal("failed to get loms address")
	}
	lomsConn, err := grpc.Dial(lomsAdd, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect loms to server: ", err)
	}
	defer lomsConn.Close()

	lomsClient := grpcLoms.New(lomsConn)
	if lomsClient == nil {
		log.Fatal("failed to init logistics and order manager service")
	}

	// product client
	productAdd := config.ConfigData.GetServiceAddress(config.Product)
	if productAdd == "" {
		log.Fatal("failed to get product address")
	}
	productConn, err := grpc.Dial(productAdd, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect product to server: ", err)
	}
	defer productConn.Close()

	productClient := grpcProduct.New(productConn)
	if productClient == nil {
		log.Fatal("failed to init product service")
	}

	/// 2. Creating server

	// создаем скелет сервера
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(interceptors.LoggingInterceptor),
		),
	)
	// для постмана: распознаем апи, используемый сервером
	reflection.Register(s)

	checkoutServer := checkoutV1.New(checkoutService, lomsClient, productClient, config.ConfigData)
	if checkoutServer == nil {
		log.Fatal("failed to init checkout server version 1")
	}
	// регистрируем поведение
	desc.RegisterCheckoutV1Server(s, checkoutServer)

	/// 3. Run

	grpcAdd := config.ConfigData.GetCommunicationAddress(config.GRPC)
	listener, err := net.Listen("tcp", grpcAdd)
	if err != nil {
		log.Fatal("failed to create listener: ", err)
	}

	log.Printf("grpc listening at %v:", listener.Addr())

	if err := s.Serve(listener); err != nil {
		log.Fatal("failed to serve: ", err)
	}
}

func initJsonRpcCommunication(checkoutService *checkout.Service, stor *storage.WrapStorage) {
	/// 1. Creating clients

	// services: loms: logistics
	lomsAdd := config.ConfigData.GetServiceAddress(config.Loms)
	logistics := jsonLoms.New(lomsAdd)
	if logistics == nil {
		log.Fatal("failed to init logistics manager service")
	}

	// services: loms: order manager
	orders := jsonLoms.New(lomsAdd)
	if orders == nil {
		log.Fatal("failed to init order manager service")
	}

	// services: product
	productAdd := config.ConfigData.GetServiceAddress(config.Product)
	product := jsonProduct.New(productAdd)
	if product == nil {
		log.Fatal("failed to init product service")
	}

	/// 2. Setting up communication using handlers

	// handlers: addToCart
	addToCartHandler := addtocart.New(checkoutService)
	http.Handle("/addToCart", server.New(addToCartHandler.Handle))

	// handlers: deleteFromCart
	deleteFromCartHandler := deletefromcart.New(checkoutService)
	http.Handle("/deleteFromCart", server.New(deleteFromCartHandler.Handle))

	// handlers: cartList
	cartListHandler := cartlist.New(checkoutService)
	http.Handle("/cartListt", server.New(cartListHandler.Handle))

	// handlers: purchase
	purchaseHandler := purchase.New(checkoutService)
	http.Handle("/purchase", server.New(purchaseHandler.Handle))

	/// 3. Run

	httpAdd := config.ConfigData.GetCommunicationAddress(config.HTTP)
	log.Println("listening http at ", httpAdd)
	err := http.ListenAndServe(httpAdd, nil)
	if err != nil {
		log.Fatal("listening and serve: ", err)
	}
	log.Fatal("unable to listen to http", err)
}
