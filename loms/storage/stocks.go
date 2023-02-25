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

// возврыщает информацию о наличии товара на разных складах в случае успеха
func (s *WrapStorage) Stocks(ctx context.Context, sku uint32) ([]model.Stock, error) {
	// 1. прошерстить все стоки в поисках товара
	storStocks, err := s.logisticsStor.CheckStocks(ctx, sku)
	if err != nil {
		return nil, errors.WithMessage(err, "checking warehouses")
	}

	stocks := s.convertToModelStocks(ctx, storStocks)

	return stocks, nil
}
