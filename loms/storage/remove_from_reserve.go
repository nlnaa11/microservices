package storage

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
)

func (s *WrapStorage) RemoveFromReserve(ctx context.Context, items []model.Item) error {
	if len(items) == 0 {
		return ErrEmptyItems
	}

	storItems := s.convertToStorItems(ctx, items)

	return s.logisticsStor.RemoveFromReserve(ctx, storItems)
}
