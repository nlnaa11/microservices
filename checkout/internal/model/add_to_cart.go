package model

import (
	"context"

	"github.com/pkg/errors"
)

var (
	ErrInsufficientStocks = errors.New("insufficient stocks")
)

type Stock struct {
	WarehouseId int64  `json:"warehouseId"`
	Count       uint64 `json:"count"`
}

func (m *Model) AddToCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	stocks, err := m.logisticsManager.Stocks(ctx, sku)
	if err != nil {
		return errors.WithMessage(err, "checking stocks")
	}

	addToCart := func() {
		item := m.createItem(sku, count)
		m.stor.AddToCart(ctx, user, *item)
	}

	cnt := int64(count)
	for _, stock := range stocks {
		cnt -= int64(stock.Count)
		if cnt <= 0 {
			addToCart()
			return nil
		}
	}

	return ErrInsufficientStocks
}

type Item struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

func (m *Model) createItem(sku uint32, count uint16) *Item {
	return &Item{
		Sku:   sku,
		Count: count,
	}
}
