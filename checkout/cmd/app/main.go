package main

import (
	"log"
	"net/http"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/config"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/handlers/addtocart"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/handlers/deletefromcart"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/handlers/listcart"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/handlers/purchase"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/model"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/services/loms"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/services/product"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/storage"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/storage/memory"
	"gitlab.ozon.dev/nlnaa/homework-1/libs/wrappers/server"
)

const (
	port = ":8080"

	ListeningHTTP      = "listening http at "
	UnableToListenHTTP = "unable to listen to http"
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

	// services: loms: logistics
	logistics := loms.New(config.ConfigData.Services.Loms, loms.PathStocks)
	if logistics == nil {
		log.Fatal("failed to init logistics manager service")
	}

	// services: loms: order manager
	orders := loms.New(config.ConfigData.Services.Loms, loms.PathCreateOrder)
	if orders == nil {
		log.Fatal("failed to init order manager service")
	}

	// services: product
	product := product.New(config.ConfigData.Services.Product, product.PathProducts)
	if product == nil {
		log.Fatal("failed to init product service")
	}

	// business logic
	checkoutModel := model.New(product, logistics, orders, stor)

	// handlers: addToCart
	addToCartHandler := addtocart.New(checkoutModel)

	http.Handle("/addToCart", server.New(addToCartHandler.Handle))

	log.Println(ListeningHTTP, port)
	err = http.ListenAndServe(port, nil)
	log.Fatal(UnableToListenHTTP, err)

	// handlers: deleteFromCart
	deleteFromCartHandler := deletefromcart.New(checkoutModel)

	http.Handle("/deleteFromCart", server.New(deleteFromCartHandler.Handle))

	log.Println(ListeningHTTP, port)
	err = http.ListenAndServe(port, nil)
	log.Fatal(UnableToListenHTTP, err)

	// handlers: listCart
	listCartHandler := listcart.New(checkoutModel)

	http.Handle("/listCart", server.New(listCartHandler.Handle))

	log.Println(ListeningHTTP, port)
	err = http.ListenAndServe(port, nil)
	log.Fatal(UnableToListenHTTP, err)

	// handlers: purchase
	purchaseHandler := purchase.New(checkoutModel)

	http.Handle("/purchase", server.New(purchaseHandler.Handle))

	log.Println(ListeningHTTP, port)
	err = http.ListenAndServe(port, nil)
	log.Fatal(UnableToListenHTTP, err)

}
