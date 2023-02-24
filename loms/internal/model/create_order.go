package model

import (
	"context"

	"github.com/pkg/errors"
)

const (
	StatusNew             = 0
	StatusFailed          = 1
	StatusAwaitingPayment = 2
)

type OrderInfo struct {
	OrderId uint64
	Status  uint16
}

func (m *Model) CreateOrder(ctx context.Context, user int64, items []Item) (OrderInfo, error) {
	order := OrderInfo{
		OrderId: 0,
		Status:  StatusNew,
	}

	// 1. создать новый заказ в хранилище, получить идентификатор заказа
	orderId, err := m.stor.CreateOrder(ctx, user, items)
	if err != nil {
		order.Status = StatusFailed
		return order, errors.WithMessage(err, "creating order")
	}

	order.OrderId = orderId

	// 2. Зарезервировать товары
	// 2.1. Удалось -- в статус awaiting payment
	// 2.2. Не удалось -- в статус failed
	_, err = m.stor.ReserveItems(ctx, items)
	if err != nil {
		order.Status = StatusFailed
		_ = m.stor.SetOrderStatus(ctx, orderId, StatusFailed)
		return order, errors.WithMessage(err, "reserve items")
	}

	err = m.stor.SetOrderStatus(ctx, orderId, StatusAwaitingPayment)
	if err != nil {
		// отменить резерв
		_ = m.stor.RemoveFromReserve(ctx, items)
		order.Status = StatusFailed
		return order, errors.WithMessage(err, "setting new status")
	}

	order.Status = StatusAwaitingPayment
	return order, nil
}
