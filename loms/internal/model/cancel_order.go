package model

import (
	"context"

	"github.com/pkg/errors"
)

const (
	StatusCancelled = 4
)

func (m *Model) CancelOrder(ctx context.Context, orderId uint64) error {
	// 1. получить заказ по id
	order, err := m.stor.GetOrderData(ctx, orderId)
	if err != nil {
		return errors.WithMessage(err, "getting order by id")
	}
	// 2. отменить резерв
	_ = m.stor.RemoveFromReserve(ctx, order.Items)

	// 3. перевести заказ в статус "отменен"
	_ = m.stor.SetOrderStatus(ctx, orderId, StatusCancelled)

	return nil
}
