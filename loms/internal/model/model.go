package model

import "context"

type Storage interface {
	/** orders **/

	CreateOrder(ctx context.Context, user int64, items []Item) (uint64, error)
	SetOrderStatus(ctx context.Context, orderId uint64, status uint16) error
	GetOrderData(ctx context.Context, orderId uint64) (OrderData, error)

	/** logistics **/

	ReserveItems(ctx context.Context, items []Item) (map[uint32]uint64, error)
	// for cancel
	RemoveFromReserve(ctx context.Context, items []Item) error
	// after payment
	RemoveItems(ctx context.Context, items []Item) error
	Stocks(ctx context.Context, sku uint32) ([]Stock, error)
}

type Model struct {
	stor Storage
}

func New(stor Storage) *Model {
	return &Model{
		stor: stor,
	}
}
