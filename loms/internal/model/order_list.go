package model

import (
	"context"

	"github.com/pkg/errors"
)

type Item struct {
	Sku   uint32
	Count uint16
}

type OrderData struct {
	Status uint16
	User   int64
	Items  []Item
}

func (m *Model) ListOfOrder(ctx context.Context, orderId uint64) (OrderData, error) {
	var order OrderData

	// 1. получить заказ из хранилища
	order, err := m.stor.GetOrderData(ctx, orderId)
	if err != nil {
		return order, errors.WithMessage(err, "getting order by id")
	}

	return order, nil
}

// type StatusType uint16

// const (
// 	StatusNew             StatusType = 0
// 	StatusFailed          StatusType = 1
// 	StatusAwaitingPayment StatusType = 2
// 	StatusPayed           StatusType = 3
// 	StatusCancelled       StatusType = 4
// )

// func (s StatusType) String() string {
// 	switch s {
// 	case StatusNew:
// 		return "new"
// 	case StatusFailed:
// 		return "failed"
// 	case StatusAwaitingPayment:
// 		return "awaiting payment"
// 	case StatusPayed:
// 		return "payed"
// 	case StatusCancelled:
// 		return "cancelled"
// 	default:
// 		return "unknown"
// 	}
// }
