package storage

import "context"

type CartStorage interface {
	AddToCart(context context.Context, item Item, user int64) error
	RemoveFromCart(ctx context.Context, item Item, user int64) error
	GetAll(ctx context.Context, user int64) (*Cart, error)
}

type WrapStorage struct {
	cartStor CartStorage
}

func New(stor CartStorage) *WrapStorage {
	return &WrapStorage{
		cartStor: stor,
	}
}
