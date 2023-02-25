package memory

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/storage"
)

func (m *MemoryWrapper) convertToItem(ctx context.Context, storItem *storage.Item) *Item {
	return &Item{
		sku:   storItem.Sku,
		count: storItem.Count,
	}
}

func (m *MemoryWrapper) convertToStorItems(ctx context.Context, items []Item) []storage.Item {
	storItems := make([]storage.Item, 0, len(items))
	for _, item := range items {
		storItems = append(storItems, storage.Item{
			Sku:   item.sku,
			Count: item.count,
		})
	}
	return storItems
}

func (m *MemoryWrapper) convertToStorCart(ctx context.Context, cart *Cart) *storage.Cart {
	storItems := m.convertToStorItems(ctx, cart.items)
	return &storage.Cart{
		Items: storItems,
	}
}
