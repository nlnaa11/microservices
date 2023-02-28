package logisticsmem

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/loms/storage"
)

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
