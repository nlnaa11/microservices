package storage

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/model"
)

type Cart struct {
	Items []Item
}

func (s *WrapStorage) GetCart(ctx context.Context, user int64) (model.Cart, error) {
	cart, err := s.cartStor.GetAll(ctx, user)
	if err != nil {
		return model.Cart{}, errors.WithMessage(err, "getting cart")
	}

	return *s.convertToModelCart(ctx, cart), nil
}
