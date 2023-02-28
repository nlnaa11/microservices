package storage

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/model"
)

type Item struct {
	Sku   uint32
	Count uint64
}

func (s *WrapStorage) AddToCart(ctx context.Context, user int64, modelItem model.Item) error {
	item := s.convertToItem(ctx, modelItem)

	if err := s.cartStor.AddToCart(ctx, item, user); err != nil {
		return errors.WithMessage(err, "adding to storage")
	}
	return nil
}
