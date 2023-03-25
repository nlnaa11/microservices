// TODO: "запихивать всё в одно место не очень хорошо, мейн не должен
// быть прям совсем чистый, это в будущем лучше переработать
// (из разряда инит клиентов отдельно, ран отдельно и в мейне вызовы)"

package main

import (
	"context"
	"log"

	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/config"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/loms"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	ctx := context.Background()

	app, err := loms.New(ctx, &config.Data)
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
