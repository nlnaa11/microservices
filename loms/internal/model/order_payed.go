package model

import (
	"context"

	"github.com/pkg/errors"
)

func (m *Model) OrderPayed(ctx context.Context, orderId uint64) error {
	// 1. получить заказ по id
	orderData, err := m.stor.GetOrderData(ctx, orderId)
	if err != nil {
		return errors.WithMessage(err, "getting order by id")
	}
	// 2. удалить товары со складов и из резерва
	err = m.stor.RemoveItems(ctx, orderData.Items)
	if err != nil {
		return errors.WithMessage(err, "removing items")
	}
	// 3. перевести заказ в статус "оплачен"
	_ = m.stor.SetOrderStatus(ctx, orderId, StatusPayed.String())

	return nil
}
