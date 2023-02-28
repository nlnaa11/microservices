package storage

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
)

// удаляет товары из резерва в случае отмены заказа или при каком-л сбое
func (s *WrapStorage) RemoveFromReserve(ctx context.Context, modelItems []model.Item) error {
	if len(modelItems) == 0 {
		return errors.ErrEmptyItems
	}

	items := s.convertToItems(ctx, modelItems)

	return s.logisticsStor.RemoveFromReserve(ctx, items)
}
