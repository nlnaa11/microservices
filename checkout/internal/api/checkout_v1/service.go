package checkout_v1

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/grpc/loms"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/grpc/product"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/model"
	desc "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/checkout_v1"
)

type Service interface {
	AddToCart(ctx context.Context, user int64, sku uint32, count uint64) error
	DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint64) error
	CartList(ctx context.Context, user int64) (model.CartList, error)
	Purchase(ctx context.Context, user int64) (model.OrderInfo, error)
}

type TokenGetter interface {
	// Токен для подключения к Продуктовому сервису
	GetToken() string
}

type Implemantation struct {
	desc.UnimplementedCheckoutV1Server

	checkoutService Service
	lomsClient      loms.Client
	productService  product.Client

	tokenGetter TokenGetter
}

func New(checkService Service, loms loms.Client, product product.Client, tokenGetter TokenGetter) *Implemantation {
	return &Implemantation{
		desc.UnimplementedCheckoutV1Server{},

		checkService,
		loms,
		product,

		tokenGetter,
	}
}
