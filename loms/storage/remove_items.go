package storage

import (
	"context"

	"github.com/pkg/errors"
	internalErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
)

// удаляет товары со складов и из резерва в случае оплаченной покупки
func (s *WrapStorage) RemoveItems(ctx context.Context, modelItems []model.Item) error {
	if len(modelItems) == 0 {
		return internalErr.ErrEmptyItems
	}

	items := s.convertToItems(ctx, modelItems)

	// 1. удалить со складов
	err := s.logisticsStor.RemoveFromWarehouses(ctx, items)
	if err != nil {
		return errors.WithMessage(err, "removing from warehouses")
	}
	// 2. удалить из резерва
	err = s.logisticsStor.RemoveFromReserve(ctx, items)
	if err != nil {
		return errors.WithMessage(err, "removing from reserve")
	}

	return nil
}
