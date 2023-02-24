package storage

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
)

func (s *WrapStorage) ReserveItems(ctx context.Context, items []model.Item) (map[uint32]uint64, error) {
	if len(items) == 0 {
		return nil, ErrEmptyItems
	}

	storItems := s.convertToStorItems(ctx, items)

	// 1. Проверить запасы на складе (доступность)
	inDeficit, err := s.checkAvailability(ctx, storItems)
	if err != nil {
		return inDeficit, errors.WithMessage(err, "checking stocks")
	}
	if len(inDeficit) > 0 {
		return inDeficit, errors.WithMessage(err, "deficits of items")
	}

	return nil, s.logisticsStor.AddToReserve(ctx, storItems)
}

func (s *WrapStorage) checkAvailability(ctx context.Context, items []Item) (map[uint32]uint64, error) {
	inDeficit := make(map[uint32]uint64)

	// 1. проверить запасы
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
