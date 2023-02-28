package ordermem

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

func (o *OrderMemoryWrapper) SetOrderStatus(ctx context.Context, orderId uint64, status string) error {
	order, ok := o.orders[orderId]
	if !ok {
		return errors.ErrOrderNotFound
	}
	if status == StatusUnknown.String() {
		return errors.ErrUnknownOrderStatus
	}

	order.status = StatusFromString(status)

	return nil
}
