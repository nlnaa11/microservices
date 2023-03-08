package loms

import (
	"context"

	lomsServiceApi "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/loms_v1"
	"google.golang.org/grpc"
)

// For compliance, it is necessary to implement the  loms.Client interface:
var _ Client = (*client)(nil)

type Client interface {
	// Проверяем запасы на наличе необходимого Количества товара с Артикулом
	Stocks(ctx context.Context, req *lomsServiceApi.StocksRequest) (*lomsServiceApi.StocksResponse, error)
	// Регистрируем новый заказ для Пользователя
	CreateOrder(ctx context.Context, req *lomsServiceApi.CreateOrderRequest) (*lomsServiceApi.CreateOrderResponse, error)
}

type client struct {
	lomsClient lomsServiceApi.LomsV1Client
}

func New(cc *grpc.ClientConn) *client {
	return &client{
		lomsClient: lomsServiceApi.NewLomsV1Client(cc),
	}
}
