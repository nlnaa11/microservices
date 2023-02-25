package logisticsmem

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/loms/storage"
)

func (l *LogisticstMemoryWrapper) AddToReserve(ctx context.Context, items []storage.Item) error {
	for _, item := range items {
		l.reserveItems[item.Sku] += item.Count
	}
	return nil
}
