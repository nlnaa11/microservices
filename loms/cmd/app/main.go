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

	//handlers: createOrder
	createOrderHandler := createorder.New(lomsModel)
	http.Handle("/createOrder", server.New(createOrderHandler.Handle))

	// // handlers: listOrder
	listOrderHandler := orderlist.New(lomsModel)
	http.Handle("/listOrder", server.New(listOrderHandler.Handle))

	// handlers: orderPayed
	orderPayedHandler := orderpayed.New(lomsModel)
	http.Handle("/orderPayed", server.New(orderPayedHandler.Handle))

	// handlers: cancelOrder
	cancelOrderHandler := cancelorder.New(lomsModel)
	http.Handle("/cancelOrder", server.New(cancelOrderHandler.Handle))

	log.Println("listening http at ", config.ConfigData.GetPort())
	err = http.ListenAndServe(config.ConfigData.GetPort(), nil)
	if err != nil {
		log.Fatal("listening and serve: ", err)
	}
	log.Fatal("unable to listen to http", err)
}
