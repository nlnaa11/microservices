package storage

import (
	"context"

	"github.com/pkg/errors"
	internalErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
)

// возвращает ошибку, если хотя бы по одному товару недостаток.
// при ошибке также возвращает возможное количество тех товаров,
// которых недостаточно для заказа.
// если всех товаров достаточно, оба возвращаемых значения nil
func (s *WrapStorage) ReserveItems(ctx context.Context, modelItems []model.Item) (map[uint32]uint64, error) {
	if len(modelItems) == 0 {
		return nil, internalErr.ErrEmptyItems
	}

	items := s.convertToItems(ctx, modelItems)

	// 1. Проверить запасы на складе (доступность)
	inDeficit, err := s.checkAvailability(ctx, items)
	if err != nil {
		return inDeficit, errors.WithMessage(err, "checking stocks")
	}
	if len(inDeficit) > 0 {
		return inDeficit, errors.WithMessage(err, "deficits of items")
	}

	return nil, s.logisticsStor.AddToReserve(ctx, items)
}

func (s *WrapStorage) checkAvailability(ctx context.Context, items []Item) (map[uint32]uint64, error) {
	inDeficit := make(map[uint32]uint64)

	for _, item := range items {
		stocks, err := s.logisticsStor.CheckStocks(ctx, item.Sku)
		if err != nil {
			inDeficit[item.Sku] = 0
			continue
		}

		var availabilityCount int64
		requiredCount := int64(item.Count)
		isEnough := false

		for _, stock := range stocks {
			requiredCount -= int64(stock.Count)
			if requiredCount <= 0 {
				isEnough = true
				break
			}
			availabilityCount += int64(stock.Count)
		}

		if !isEnough {
			inDeficit[item.Sku] = uint64(availabilityCount)
		}
	}

	return inDeficit, nil
}
