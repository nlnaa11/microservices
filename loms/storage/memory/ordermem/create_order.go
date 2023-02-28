package ordermem

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/loms/storage"
)

func (o *OrderMemoryWrapper) CreateOrder(ctx context.Context, user int64, storItems []storage.Item) (uint64, error) {
	orderId := o.nextOrderId
	o.nextOrderId++

	items := o.convertToItems(ctx, storItems)

	orderData := OrderData{
		status: StatusNew,
		user:   user,
		items:  items,
	}

	// если что, данные полностью перезапишутся. что ok
	o.orders[orderId] = &orderData

	return orderId, nil
}
