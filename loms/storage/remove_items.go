package storage

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
)

func (s *WrapStorage) RemoveItems(ctx context.Context, items []model.Item) error {
	if len(items) == 0 {
		return ErrEmptyItems
	}

	storItems := s.convertToStorItems(ctx, items)

	// 1. удалить со складов
	err := s.logisticsStor.RemoveFromWarehouses(ctx, storItems)
	if err != nil {
		return errors.WithMessage(err, "removing from warehouses")
	}
	// 2. удалить из резерва
	err = s.logisticsStor.RemoveFromReserve(ctx, storItems)
	if err != nil {
		return errors.WithMessage(err, "removing from reserve")
	}

	return nil
}
