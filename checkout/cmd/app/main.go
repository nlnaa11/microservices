// TODO: "запихивать всё в одно место не очень хорошо, мейн не должен
// быть прям совсем чистый, это в будущем лучше переработать
// (из разряда инит клиентов отдельно, ран отдельно и в мейне вызовы)"

package main

import (
	"context"
	"log"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/checkout"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/config"
)

func main() {
	// config
	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	ctx := context.Background()

	app, err := checkout.New(ctx, &config.Data)
	if err != nil {
		log.Fatalf("Failed to create app: %s\n", err.Error())
	}
	if app == nil {
		log.Fatal("Failed to create app")
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("Failed to run app: %s\n", err.Error())
	}
}
