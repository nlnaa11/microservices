// только не бей!

package storage

import (
	"context"
)

type LogisticsStorage interface {
	CheckStocks(ctx context.Context, sku uint32) ([]Stock, error)
	AddToReserve(ctx context.Context, items []Item) error
	RemoveFromReserve(ctx context.Context, items []Item) error
	RemoveFromWarehouses(ctx context.Context, items []Item) error
}

type OrderStorage interface {
	CreateOrder(ctx context.Context, user int64, items []Item) (uint64, error)
	SetOrderStatus(ctx context.Context, orderId uint64, status string) error
	// информация о coставе заказа и его статусе
	GetOrderData(ctx context.Context, orderId uint64) (*OrderData, error)
}

type WrapStorage struct {
	logisticsStor LogisticsStorage
	orderStor     OrderStorage
}

func New(logStor LogisticsStorage, orderStor OrderStorage) *WrapStorage {
	return &WrapStorage{
		logisticsStor: logStor,
		orderStor:     orderStor,
	}
}
