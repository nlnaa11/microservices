package storage

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

func (s *WrapStorage) SetOrderStatus(ctx context.Context, orderId uint64, status string) error {
	if status == StatusUnknown {
		return errors.ErrUnknownOrderStatus
	}

	// проверка сущестования заказа внутри
	return s.orderStor.SetOrderStatus(ctx, orderId, status)
}
