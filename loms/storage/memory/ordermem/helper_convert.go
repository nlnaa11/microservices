package ordermem

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/loms/storage"
)

func (o *OrderMemoryWrapper) convertToItems(ctx context.Context, storItems []storage.Item) []Item {
	items := make([]Item, 0, len(storItems))
	for _, storItem := range storItems {
		items = append(items, Item{
			sku:   storItem.Sku,
			count: storItem.Count,
		})
	}
	return items
}

func (o *OrderMemoryWrapper) convertToStorItems(ctx context.Context, items []Item) []storage.Item {
	storItems := make([]storage.Item, 0, len(items))
	for _, item := range items {
		storItems = append(storItems, storage.Item{
			Sku:   item.sku,
			Count: item.count,
		})
	}
	return storItems
}

func (o *OrderMemoryWrapper) convertToStorOrderData(ctx context.Context, orderData *OrderData) *storage.OrderData {
	storItems := o.convertToStorItems(ctx, orderData.items)

	storOrderData := storage.OrderData{
		Status: orderData.status.String(),
		User:   orderData.user,
		Items:  storItems,
	}
	return &storOrderData
}
