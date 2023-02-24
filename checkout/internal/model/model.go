package model

import "context"

type LogisticsManager interface {
	Stocks(ctx context.Context, sku uint32) ([]Stock, error)
}

type OrderManager interface {
	CreateOrder(ctx context.Context, user int64, items []Item) (*Order, error)
}

type ProductInformator interface {
	GetProductInfo(ctx context.Context, token string, sku uint32) (ItemInfo, error)
}

type Storage interface {
	AddToCart(ctx context.Context, user int64, item Item) error
	RemoveFromCart(ctx context.Context, user int64, item Item) error
	GetCart(ctx context.Context, user int64) (Cart, error)
}

type TokenGetter interface {
	GetToken() string
}

type Model struct {
	productInformator ProductInformator
	logisticsManager  LogisticsManager
	orderManager      OrderManager
	stor              Storage
	tokenGetter       TokenGetter
}

func New(prodInformator ProductInformator, logManager LogisticsManager,
	ordManager OrderManager, stor Storage, tokenGetter TokenGetter) *Model {
	return &Model{
		productInformator: prodInformator,
		logisticsManager:  logManager,
		orderManager:      ordManager,
		stor:              stor,
		tokenGetter:       tokenGetter,
	}
}
