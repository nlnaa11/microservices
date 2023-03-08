package product

import (
	"context"
	"fmt"

	productServiceApi "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/product"
)

func (c *client) GetProduct(ctx context.Context, req *productServiceApi.GetProductRequest) (*productServiceApi.GetProductResponse, error) {
	fmt.Printf("GetProduct: getting information about the %d product\n", req.GetSku())

	productInfo, err := c.productClient.GetProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	return productInfo, nil
}
