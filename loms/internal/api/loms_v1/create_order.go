package loms_v1

import (
	"context"

	"github.com/pkg/errors"
	desc "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/loms_v1"
	converter "gitlab.ozon.dev/nlnaa/homework-1/loms/internal/converter/loms_v1"
)

func (i *Implemantation) CreateOrder(ctx context.Context, req *desc.CreateOrderRequest) (*desc.CreateOrderResponse, error) {
	items := converter.ToItems(req.GetItems())

	order, err := i.lomsService.CreateOrder(ctx, req.GetUser(), items)
	if err != nil {
		return nil, errors.WithMessage(err, "getting order list")
	}

	return &desc.CreateOrderResponse{
		OrderId: int64(order.OrderId),
		Status:  order.Status,
	}, nil
}
