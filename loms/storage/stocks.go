package storage

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
)

type Stock struct {
	WarehouseId int64
	Count       uint64
}

func (s *WrapStorage) Stocks(ctx context.Context, sku uint32) ([]model.Stock, error) {
	// 1. прошерстить все стоки в поисках товара
	storStocks, err := s.logisticsStor.CheckStocks(ctx, sku)
	if err != nil {
		return nil, errors.WithMessage(err, "checking warehouses")
	}

	stocks := s.convertToStocks(ctx, storStocks)

	return stocks, nil
}

func (s *WrapStorage) convertToStocks(ctx context.Context, storStocks []Stock) []model.Stock {
	stocks := make([]model.Stock, 0, len(storStocks))

	for _, storStock := range storStocks {
		stocks = append(stocks, model.Stock{
			WarehouseId: storStock.WarehouseId,
			Count:       storStock.Count,
		})
	}

	return stocks
}
