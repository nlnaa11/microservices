package checkout

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/model"
)

func (s *Service) CartList(ctx context.Context, user int64) (model.CartList, error) {
	fmt.Printf("A cart list of the %d user\n", user)

	itemsCount := gofakeit.IntRange(1, 11)
	items := make([]model.ItemInfo, 0, itemsCount)
	var totalPrice float64

	for i := 0; i < itemsCount; i++ {
		items = append(items, model.ItemInfo{
			Sku:   uint32(gofakeit.IntRange(1076963, 149097564)),
			Count: uint64(gofakeit.IntRange(1, 32)),
			Name:  gofakeit.CarMaker(),
			Price: gofakeit.Float64Range(1, 200000),
		})

		totalPrice += items[i].Price
	}

	fmt.Printf("%d user has %d types of items in his cart for a total price of %v\n",
		user, itemsCount, totalPrice)

	return model.CartList{
		ItemsInfo:  items,
		TotalPrice: totalPrice,
	}, nil
}
