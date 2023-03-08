package product

import (
	"context"

	productServiceApi "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/product"
	"google.golang.org/grpc"
)

// For compliance, it is necessary to implement the product.Client interface:
var _ Client = (*client)(nil)

type Client interface {
	// Получаем Имя и Цену товара по Артикулу
	GetProduct(ctx context.Context, in *productServiceApi.GetProductRequest) (*productServiceApi.GetProductResponse, error)
}

type client struct {
	productClient productServiceApi.ProductServiceClient
}

func New(cc *grpc.ClientConn) *client {
	return &client{
		productClient: productServiceApi.NewProductServiceClient(cc),
	}
}
