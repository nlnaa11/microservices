package model

import (
	"context"

	"github.com/pkg/errors"
	internalErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

type Stock struct {
	WarehouseId int64  `json:"warehouseId"`
	Count       uint64 `json:"count"`
}

func (m *Model) AddToCart(ctx context.Context, user int64, sku uint32, count uint64) error {
	stocks, err := m.logisticsManager.Stocks(ctx, sku)
	if err != nil {
		return errors.WithMessage(err, "checking stocks")
	}

	addToCart := func() {
		item := Item{
			Sku:   sku,
			Count: count,
		}
		_ = m.stor.AddToCart(ctx, user, item)
	}

	// если в запасе товара не меньше, чем требуется, добавляем в корзину
	cnt := int64(count)
	for _, stock := range stocks {
		cnt -= int64(stock.Count)
		if cnt <= 0 {
			addToCart()
			return nil
		}
	}

	return internalErr.ErrInsufficientStocks
}
