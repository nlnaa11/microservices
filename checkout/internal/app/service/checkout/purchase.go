package checkout

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/model"
	internalErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

func (s *Service) Purchase(ctx context.Context, user int64) (model.Cart, error) {
	fmt.Printf("%d user makes a purchase\n", user)

	var cart model.Cart

	// 1. получить корзину
	cart, err := s.cartsRepo.GetCart(ctx, user)
	if err != nil {
		return cart, errors.WithMessage(err, "getting cart")
	}
	if len(cart.Items) == 0 {
		return cart, internalErr.ErrEmptyCart
	}

	return cart, nil
}
