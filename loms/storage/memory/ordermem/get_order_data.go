package ordermem

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/storage"
)

func (o *OrderMemoryWrapper) GetOrderData(ctx context.Context, orderId uint64) (*storage.OrderData, error) {
	orderData, ok := o.orders[orderId]
	if !ok {
		return nil, errors.ErrOrderNotFound
	}

	return o.convertToStorOrderData(ctx, orderData), nil
}
