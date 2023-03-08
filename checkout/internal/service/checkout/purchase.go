package checkout

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/model"
)

func (s *Service) Purchase(ctx context.Context, user int64) (model.OrderInfo, error) {
	fmt.Printf("%d user makes a purchase\n", user)

	// 1. получить корзину
	cart, _ := s.CartList(ctx, user)

	fmt.Printf("%d user has %d types of items in his cart for a total price of %v\n",
		user, len(cart.ItemsInfo), cart.TotalPrice)

	return model.OrderInfo{
		OrderId: int64(gofakeit.IntRange(1, 100)),
		Status:  model.AwaitingPayment,
	}, nil
}
