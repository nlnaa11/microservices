package logisticsmem

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/loms/storage"
)

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
