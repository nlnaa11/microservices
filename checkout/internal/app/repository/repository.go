package repository

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/model"
)

type CartsRepository interface {
	AddToCart(ctx context.Context, user int64, item model.Item) error
	RemoveFromCart(ctx context.Context, user int64, item model.Item) error
	GetCart(cxt context.Context, user int64) (model.Cart, error)
	ClearCart(ctx context.Context, user int64) error
}
