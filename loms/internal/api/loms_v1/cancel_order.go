package loms_v1

import (
	"context"

	"github.com/pkg/errors"
	desc "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/loms_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implemantation) CancelOrder(ctx context.Context, req *desc.CancelOrderRequest) (*emptypb.Empty, error) {
	err := i.lomsService.CancelOrder(ctx, req.GetOrderId())
	if err != nil {
		return &emptypb.Empty{}, errors.WithMessage(err, "cancel order")
	}

	return &emptypb.Empty{}, nil
}
