package model

import (
	"context"

	"github.com/pkg/errors"
)

type OrderInfo struct {
	OrderId uint64
	Status  string
}

func (m *Model) CreateOrder(ctx context.Context, user int64, items []Item) (OrderInfo, error) {
	order := OrderInfo{
		OrderId: 0,
		Status:  StatusNew.String(),
	}

	// 1. создать новый заказ в хранилище, получить идентификатор заказа
	orderId, err := m.stor.CreateOrder(ctx, user, items)
	if err != nil {
		order.Status = StatusFailed.String()
		return order, errors.WithMessage(err, "creating order")
	}

	order.OrderId = orderId

	// 2. Зарезервировать товары
	// 2.1. Удалось -- в статус awaiting payment
	// 2.2. Не удалось -- в статус failed
	_, err = m.stor.ReserveItems(ctx, items)
	if err != nil {
		order.Status = StatusFailed.String()
		_ = m.stor.SetOrderStatus(ctx, orderId, StatusFailed.String())
		return order, errors.WithMessage(err, "reserve items")
	}

	err = m.stor.SetOrderStatus(ctx, orderId, StatusAwaitingPayment.String())
	if err != nil {
		// отменить резерв
		_ = m.stor.RemoveFromReserve(ctx, items)
		order.Status = StatusFailed.String()
		return order, errors.WithMessage(err, "setting new status")
	}

	order.Status = StatusAwaitingPayment.String()
	return order, nil
}
