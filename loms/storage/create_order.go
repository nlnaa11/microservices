package storage

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
)

type Item struct {
	Sku   uint32
	Count uint64
}

// возвразает валидный orderId в случае успеха
func (s *WrapStorage) CreateOrder(ctx context.Context, user int64, modelItems []model.Item) (uint64, error) {
	if len(modelItems) == 0 {
		return 0, errors.ErrEmptyItems
	}

	items := s.convertToItems(ctx, modelItems)

	orderId, err := s.orderStor.CreateOrder(ctx, user, items)

	return orderId, err
}
