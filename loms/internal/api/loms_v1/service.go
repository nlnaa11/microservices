package loms_v1

import (
	"context"

	desc "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/loms_v1"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
)

type Service interface {
	CancelOrder(ctx context.Context, orderId uint64) error
	CreateOrder(ctx context.Context, user int64, items []model.Item) (model.Order, error)
	OrderList(ctx context.Context, orderId uint64) (model.OrderInfo, error)
	OrderPayed(ctx context.Context, orderId uint64) error
	Stocks(ctx context.Context, sku uint32) ([]model.Stock, error)
}

type Implemantation struct {
	desc.UnimplementedLomsV1Server

	lomsService Service
}

func New(loms Service) *Implemantation {
	return &Implemantation{
		desc.UnimplementedLomsV1Server{},

		loms,
	}
}
