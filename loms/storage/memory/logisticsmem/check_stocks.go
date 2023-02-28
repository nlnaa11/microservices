package logisticsmem

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/storage"
)

func (l *LogisticstMemoryWrapper) CheckStocks(ctx context.Context, sku uint32) ([]storage.Stock, error) {
	stocks := make([]storage.Stock, 0, len(l.warehouses))

	for id, warehouse := range l.warehouses {
		if totalCount, ok := warehouse.items[sku]; ok {
			count := int64(totalCount)

			if reserveCount, ok := l.reserveItems[sku]; ok {
				count -= int64(reserveCount)
				if count <= 0 {
					continue
				}
			}

			stocks = append(stocks, storage.Stock{
				WarehouseId: id,
				Count:       uint64(count),
			})
		}
	}

	if len(stocks) == 0 {
		return nil, errors.ErrSkuNotFound
	}

	return stocks, nil
}
