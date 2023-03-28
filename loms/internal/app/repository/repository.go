package repository

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/model"
)

type StocksRepository interface {
	CheckStocks(ctx context.Context, sku uint32) ([]model.Stock, error)
	CheckReserve(ctx context.Context, sku uint32) (uint64, error)

	AddToReserve(ctx context.Context, items []model.Item) error
	RemoveFromReserve(ctx context.Context, items []model.Item) error
	// склады для удаления и обновления определяются в бизнес логике
	RemoveFromStocks(ctx context.Context, itemId uint32, whindicesToRemove []int64, whToUpdate *model.Stock) error
}

type OrdersRepository interface {
	CreateOrder(ctx context.Context, user int64, items []model.Item) (uint64, error)
	SetOrderStatus(ctx context.Context, orderId uint64, status string) error
	// информация о coставе заказа и его статусе
	GetOrderData(ctx context.Context, orderId uint64) (*model.OrderInfo, error)
	// информация только о составе заказа
	GetOrderItems(ctx context.Context, orderId uint64) ([]model.Item, error)
	// информация о статусе заказа
	GetOrderInfo(ctx context.Context, orderId uint64) (*model.Order, error)
}
