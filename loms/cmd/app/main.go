package main

import (
	"log"
	"net"
	"net/http"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	desc "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/loms_v1"
	"gitlab.ozon.dev/nlnaa/homework-1/libs/wrappers/server"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/api/loms_v0/cancelorder"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/api/loms_v0/createorder"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/api/loms_v0/orderlist"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/api/loms_v0/orderpayed"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/api/loms_v0/stocks"
	lomsV1 "gitlab.ozon.dev/nlnaa/homework-1/loms/internal/api/loms_v1"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/config"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/interceptors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/service/loms"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/storage"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/storage/memory/logisticsmem"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/storage/memory/ordermem"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// storages
	logisticsMem, err := logisticsmem.Init()
	if err != nil {
		log.Fatal("failed to access logistics storage", err)
	}
	ordersMem, err := ordermem.Init()
	if err != nil {
		log.Fatal("failed to access orders storage", err)
	}
	stor := storage.New(logisticsMem, ordersMem)

	// config
	err = config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	// business logic with behavior
	lomsService := loms.New(stor)
	if lomsService == nil {
		log.Fatal("failed to init loms service")
	}

	// grpc communication
	initGrpcCommunication(lomsService, stor)

	// параллельно не работает, пока не используем горутины и контексты
	// пусть будет так, чтоб не переделывать по тысячу раз
	// я починю, обещаю!

	// json rpc communication
	initJsonRpcCommunication(lomsService, stor)
}

func initGrpcCommunication(lomsService *loms.Service, stor *storage.WrapStorage) {
	/// 1. No clients

	/// 2. Creating server

	// создаем скелет сервера
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(interceptors.LoggingInterceptor),
		),
	)
	// для постмана: распознаем апи, используемый сервером
	reflection.Register(s)

	lomsServer := lomsV1.New(lomsService)
	if lomsServer == nil {
		log.Fatal("failed to init loms server version 1")
	}
	// регистрируем поведение
	desc.RegisterLomsV1Server(s, lomsServer)

	/// 3. Run

	// listener for checkout client
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

func initJsonRpcCommunication(lomsService *loms.Service, stor *storage.WrapStorage) {
	/// 1. No clients

	/// 2. Setting up communication using handlers

	// handlers: stocks
	stocksHandler := stocks.New(lomsService)
	http.Handle("/stocks", server.New(stocksHandler.Handle))

	//handlers: createOrder
	createOrderHandler := createorder.New(lomsService)
	http.Handle("/createOrder", server.New(createOrderHandler.Handle))

	// // handlers: listOrder
	listOrderHandler := orderlist.New(lomsService)
	http.Handle("/listOrder", server.New(listOrderHandler.Handle))

	// handlers: orderPayed
	orderPayedHandler := orderpayed.New(lomsService)
	http.Handle("/orderPayed", server.New(orderPayedHandler.Handle))

	// handlers: cancelOrder
	cancelOrderHandler := cancelorder.New(lomsService)
	http.Handle("/cancelOrder", server.New(cancelOrderHandler.Handle))

	/// 3. Run

	httpAdd := config.ConfigData.GetCommunicationAddress(config.HTTP)
	log.Println("listening http at ", httpAdd)
	err := http.ListenAndServe(httpAdd, nil)
	if err != nil {
		log.Fatal("listening and serve: ", err)
	}
	log.Fatal("unable to listen to http", err)
}
