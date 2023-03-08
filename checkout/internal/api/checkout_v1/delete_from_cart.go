package checkout_v1

import (
	"context"

	"github.com/pkg/errors"
	desc "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/checkout_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implemantation) DeleteFromCart(ctx context.Context, req *desc.DeleteFromCartRequest) (*emptypb.Empty, error) {
	err := i.checkoutService.DeleteFromCart(ctx, req.GetUser(), req.GetItem().GetSku(), req.GetItem().GetCount())
	if err != nil {
		return nil, errors.WithMessage(err, "deleting from cart")
	}

	return &emptypb.Empty{}, nil
}
