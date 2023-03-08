package checkout

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/model"
)

type Storage interface {
	AddToCart(ctx context.Context, user int64, item model.Item) error
	RemoveFromCart(ctx context.Context, user int64, item model.Item) error
	GetCart(ctx context.Context, user int64) (model.Cart, error)
}

type Service struct {
	stor Storage
}

func New(stor Storage) *Service {
	return &Service{
		stor: stor,
	}
}
