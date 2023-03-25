package checkout_v1

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/model"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/grpc/loms"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/grpc/product"
	desc "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/checkout_v1"
)

type Service interface {
	AddToCart(ctx context.Context, user int64, sku uint32, count uint64) error
	DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint64) error
	CartList(ctx context.Context, user int64) (model.CartList, error)
	Purchase(ctx context.Context, user int64) (model.Cart, error)
}

type Implementation struct {
	desc.UnimplementedCheckoutV1Server

	checkoutService Service
	lomsClient      loms.Client
	productService  product.Client
}

func New(checkService Service, loms loms.Client, product product.Client) *Implementation {
	return &Implementation{
		desc.UnimplementedCheckoutV1Server{},

		checkService,
		loms,
		product,
	}
}
