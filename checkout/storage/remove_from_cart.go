package storage

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/model"
)

func (s *WrapStorage) RemoveFromCart(ctx context.Context, user int64, modelItem model.Item) error {
	item := s.convertToItem(ctx, modelItem)

	if err := s.cartStor.RemoveFromCart(ctx, item, user); err != nil {
		return errors.WithMessage(err, "removing from storage")
	}
	return nil
}
