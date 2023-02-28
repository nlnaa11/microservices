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
	logistics := loms.New(config.ConfigData.Services.Loms)
	if logistics == nil {
		log.Fatal("failed to init logistics manager service")
	}

	// services: loms: order manager
	orders := loms.New(config.ConfigData.Services.Loms)
	if orders == nil {
		log.Fatal("failed to init order manager service")
	}

	// services: product
	product := product.New(config.ConfigData.Services.Product)
	if product == nil {
		log.Fatal("failed to init product service")
	}

	// business logic
	checkoutModel := model.New(product, logistics, orders, stor, config.ConfigData)

	// handlers: addToCart
	addToCartHandler := addtocart.New(checkoutModel)
	http.Handle("/addToCart", server.New(addToCartHandler.Handle))

	// handlers: deleteFromCart
	deleteFromCartHandler := deletefromcart.New(checkoutModel)
	http.Handle("/deleteFromCart", server.New(deleteFromCartHandler.Handle))

	// handlers: listCart
	listCartHandler := listcart.New(checkoutModel)
	http.Handle("/listCart", server.New(listCartHandler.Handle))

	// handlers: purchase
	purchaseHandler := purchase.New(checkoutModel)
	http.Handle("/purchase", server.New(purchaseHandler.Handle))

	log.Println("listening http at ", config.ConfigData.GetPort())
	err = http.ListenAndServe(config.ConfigData.GetPort(), nil)
	if err != nil {
		log.Fatal("listening and serve: ", err)
	}
	log.Fatal("unable to listen to http", err)
}
