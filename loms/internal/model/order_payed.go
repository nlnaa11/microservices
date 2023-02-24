package model

import (
	"context"

	"github.com/pkg/errors"
)

const (
	StatusPayed = 3
)

func (m *Model) OrderPayed(ctx context.Context, orderId uint64) error {
	// 1. получить заказ по id
	order, err := m.stor.GetOrderData(ctx, orderId)
	if err != nil {
		return errors.WithMessage(err, "getting order by id")
	}
	// 2. удалить товары со складов и из резерва
	err = m.stor.RemoveItems(ctx, order.Items)
	if err != nil {
		return errors.WithMessage(err, "removing items")
	}
	// 3. перевести заказ в статус "оплачен"
	_ = m.stor.SetOrderStatus(ctx, orderId, StatusPayed)

	return nil
}
