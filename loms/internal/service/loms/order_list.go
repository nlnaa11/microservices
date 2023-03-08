package loms

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
)

func (s *Service) OrderList(ctx context.Context, orderId uint64) (model.OrderInfo, error) {
	fmt.Printf("Getting an order list by id: %d", orderId)

	itemsCount := gofakeit.IntRange(0, 11)
	items := make([]model.Item, 0, itemsCount)
	for i := 0; i < itemsCount; i++ {
		items = append(items, model.Item{
			Sku:   gofakeit.Uint32(),
			Count: gofakeit.Uint64(),
		})
	}

	return model.OrderInfo{
		Order: model.Order{
			OrderId: orderId,
			Status:  model.StatusAwaitingPayment.String(),
		},
		User:  gofakeit.Int64(),
		Items: items,
	}, nil
}
