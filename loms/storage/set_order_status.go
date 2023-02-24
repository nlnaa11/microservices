package storage

import (
	"context"

	"github.com/pkg/errors"
)

var (
	ErrUnknownOrderStatus = errors.New("unknown order status")
)

const (
	StatusNew             = 0
	StatusFailed          = 1
	StatusAwaitingPayment = 2
	StatusPayed           = 3
	StatusCancelled       = 4
	// вот это не очень
	StatusCount = 5
)

func (s *WrapStorage) SetOrderStatus(ctx context.Context, orderId uint64, status uint16) error {
	if status >= StatusCount {
		return ErrUnknownOrderStatus
	}

	// проверка сущестования заказа внутри
	return s.orderStor.SetOrderStatus(ctx, orderId, status)
}
