package loms_v1

import (
	"context"

	"github.com/pkg/errors"
	desc "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/loms_v1"
	converter "gitlab.ozon.dev/nlnaa/homework-1/loms/internal/converter/loms_v1"
)

func (i *Implemantation) OrderList(ctx context.Context, req *desc.OrderListRequest) (*desc.OrderListResponse, error) {
	orderInfo, err := i.lomsService.OrderList(ctx, req.GetOrderId())
	if err != nil {
		return nil, errors.WithMessage(err, "getting order list")
	}

	return &desc.OrderListResponse{
		OrderInfo: converter.ToDescOrderInfo(&orderInfo),
	}, nil
}
