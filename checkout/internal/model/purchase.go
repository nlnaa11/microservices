package model

import (
	"context"

	"github.com/pkg/errors"
)

const (
	staNew             = 0
	staAwaitingPayment = 1
	staFailed          = 2
	staPayed           = 3
	stacanceled        = 4
)

type Order struct {
	OrderId int64
	Status  uint16
}

func (m *Model) Purchase(ctx context.Context, user int64) (*Order, error) {
	outOrder := Order{
		OrderId: -1,
		Status:  staNew,
	}

	// 1. получить корзину
	cart, err := m.stor.GetCart(ctx, user)
	if err != nil {
		return &outOrder, errors.WithMessage(err, "getting cart")
	}
	if len(cart.Items) < 1 {
		return &outOrder, ErrEmpryCart
	}

	// 2. создать заказ со статусом
	order, err := m.orderManager.CreateOrder(ctx, user, cart.Items)
	if err != nil {
		return &outOrder, errors.WithMessage(err, "creating order")
	}

	outOrder.OrderId = order.OrderId
	outOrder.Status = order.Status

	return &outOrder, nil
}
