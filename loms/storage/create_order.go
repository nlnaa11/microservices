package storage

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
)

var (
	ErrEmptyItems = errors.New("items are empty")
)

type Item struct {
	Sku   uint32
	Count uint64
}

func (s *WrapStorage) CreateOrder(ctx context.Context, user int64, items []model.Item) (uint64, error) {
	if len(items) == 0 {
		return 0, ErrEmptyItems
	}

	storItems := s.convertToStorItems(ctx, items)

	orderId, err := s.orderStor.CreateOrder(ctx, user, storItems)

	return orderId, err
}

func (s *WrapStorage) convertToStorItems(ctx context.Context, items []model.Item) []Item {
	storItems := make([]Item, 0, len(items))

	for _, item := range items {
		storItems = append(storItems, Item{
			Sku:   item.Sku,
			Count: uint64(item.Count),
		})
	}

	return storItems
}
