package model

import (
	"context"

	"github.com/pkg/errors"
)

type Item struct {
	Sku   uint32
	Count uint64
}

type OrderData struct {
	Status Status
	User   int64
	Items  []Item
}

func (m *Model) ListOfOrder(ctx context.Context, orderId uint64) (OrderData, error) {
	// 1. получить заказ из хранилища
	orderData, err := m.stor.GetOrderData(ctx, orderId)
	if err != nil {
		return OrderData{}, errors.WithMessage(err, "getting order by id")
	}

	return orderData, nil
}
