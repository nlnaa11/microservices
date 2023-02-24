// Это просто черновик, который помогает мне понять,
// как перемещаются данные.

package logisticsmem

import (
	"context"
	"errors"

	"gitlab.ozon.dev/nlnaa/homework-1/loms/storage"
)

var (
	ErrSkuNotFound = errors.New("sku not found")
)

const FirstWarehouseId int64 = 1

type Warehouse struct {
	// key: sku, value: count
	items map[uint32]uint64
}

type LogisticstMemoryWrapper struct {
	// key: warehouseId
	warehouses map[int64]*Warehouse
	// key: sku, value: count
	reserveItems    map[uint32]uint64
	nextWarehouseId int64
}

func Init() (*LogisticstMemoryWrapper, error) {
	return &LogisticstMemoryWrapper{
		warehouses:      make(map[int64]*Warehouse),
		reserveItems:    make(map[uint32]uint64),
		nextWarehouseId: FirstWarehouseId,
	}, nil
}

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
		return nil, ErrSkuNotFound
	}

	return stocks, nil
}

func (l *LogisticstMemoryWrapper) AddToReserve(ctx context.Context, items []storage.Item) error {
	for _, item := range items {
		l.reserveItems[item.Sku] += item.Count
	}
	return nil
}

func (l *LogisticstMemoryWrapper) RemoveFromReserve(ctx context.Context, items []storage.Item) error {
	var count int64

	for _, item := range items {
		if reserveCount, ok := l.reserveItems[item.Sku]; ok {
			count = int64(reserveCount)
			count -= int64(item.Count)
			if count <= 0 {
				l.reserveItems[item.Sku] = 0
			} else {
				l.reserveItems[item.Sku] = uint64(count)
			}
		}
	}

	return nil
}

func (l *LogisticstMemoryWrapper) RemoveFromWarehouses(ctx context.Context, items []storage.Item) error {
	var count uint64

	for _, item := range items {
		count = item.Count
		for _, warehouse := range l.warehouses {
			if totalCount, ok := warehouse.items[item.Sku]; ok {
				if totalCount <= count {
					count -= totalCount
					warehouse.items[item.Sku] = 0
					continue
				}
				warehouse.items[item.Sku] -= count
				count = 0
			}
			if count == 0 {
				break
			}
		}
	}

	return nil
}
