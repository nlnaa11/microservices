package loms_v1

import (
	"context"

	"github.com/pkg/errors"
	desc "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/loms_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implemantation) OrderPayed(ctx context.Context, req *desc.OrderPayedRequest) (*emptypb.Empty, error) {
	err := i.lomsService.OrderPayed(ctx, req.GetOrderId())
	if err != nil {
		return &emptypb.Empty{}, errors.WithMessage(err, "setting the order as paid")
	}

	return &emptypb.Empty{}, nil
}
