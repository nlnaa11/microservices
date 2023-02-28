package model

import (
	"context"

	"github.com/pkg/errors"
	internalErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

const (
	invalidOrderId = 0
)

type OrderInfo struct {
	OrderId int64
	Status  string
}

func (m *Model) Purchase(ctx context.Context, user int64) (OrderInfo, error) {
	orderInfo := OrderInfo{
		OrderId: invalidOrderId,
		Status:  Unknown,
	}

	// 1. получить корзину
	cart, err := m.stor.GetCart(ctx, user)
	if err != nil {
		return orderInfo, errors.WithMessage(err, "getting cart")
	}
	if len(cart.Items) < 1 {
		return orderInfo, internalErr.ErrEmptyCart
	}

	// 2. создать заказ со статусом
	order, err := m.orderManager.CreateOrder(ctx, user, cart.Items)
	if err != nil {
		return orderInfo, errors.WithMessage(err, "creating order")
	}

	orderInfo.OrderId = order.OrderId
	orderInfo.Status = order.Status

	return orderInfo, nil
}
