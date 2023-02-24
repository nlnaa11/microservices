package main

import (
	"log"
	"net/http"

	"gitlab.ozon.dev/nlnaa/homework-1/libs/wrappers/server"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/config"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/handlers/cancelorder"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/handlers/createorder"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/handlers/orderlist"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/handlers/orderpayed"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/handlers/stocks"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/storage"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/storage/memory/logisticsmem"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/storage/memory/ordermem"
)

const (
	port = ":8081"

	ListeningHTTP      = "listening http at "
	UnableToListenHTTP = "unable to listen to http"
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

	// there are no services

	// business logic
	lomsModel := model.New(stor)

	// handlers: stocks
	stocksHandler := stocks.New(lomsModel)

	http.Handle("/stocks", server.New(stocksHandler.Handle))

	log.Println(ListeningHTTP, port)
	err = http.ListenAndServe(port, nil)
	log.Fatal(UnableToListenHTTP, err)

	//handlers: createOrder
	createOrderHandler := createorder.New(lomsModel)

	http.Handle("/createOrder", server.New(createOrderHandler.Handle))

	log.Println(ListeningHTTP, port)
	err = http.ListenAndServe(port, nil)
	log.Fatal(UnableToListenHTTP, err)

	// // handlers: listOrder
	listOrderHandler := orderlist.New(lomsModel)

	http.Handle("/listOrder", server.New(listOrderHandler.Handle))

	log.Println(ListeningHTTP, port)
	err = http.ListenAndServe(port, nil)
	log.Fatal(UnableToListenHTTP, err)

	// handlers: orderPayed
	orderPayedHandler := orderpayed.New(lomsModel)

	http.Handle("/orderPayed", server.New(orderPayedHandler.Handle))

	log.Println(ListeningHTTP, port)
	err = http.ListenAndServe(port, nil)
	log.Fatal(UnableToListenHTTP, err)

	// handlers: cancelOrder
	cancelOrderHandler := cancelorder.New(lomsModel)

	http.Handle("/cancelOrder", server.New(cancelOrderHandler.Handle))

	log.Println(ListeningHTTP, port)
	err = http.ListenAndServe(port, nil)
	log.Fatal(UnableToListenHTTP, err)
}
