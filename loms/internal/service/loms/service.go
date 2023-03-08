package loms

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
)

type Storage interface {
	/** orders **/

	CreateOrder(ctx context.Context, user int64, items []model.Item) (uint64, error)
	SetOrderStatus(ctx context.Context, orderId uint64, status string) error
	GetOrderData(ctx context.Context, orderId uint64) (model.OrderInfo, error)

	/** logistics **/

	ReserveItems(ctx context.Context, items []model.Item) (map[uint32]uint64, error)
	// for cancel
	RemoveFromReserve(ctx context.Context, items []model.Item) error
	// after payment
	RemoveItems(ctx context.Context, items []model.Item) error
	Stocks(ctx context.Context, sku uint32) ([]model.Stock, error)
}

type Service struct {
	stor Storage
}

func New(stor Storage) *Service {
	return &Service{
		stor: stor,
	}
}
