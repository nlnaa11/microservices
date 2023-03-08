package loms

import (
	"context"
	"fmt"

	lomsServiceApi "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/loms_v1"
)

func (c *client) CreateOrder(ctx context.Context, req *lomsServiceApi.CreateOrderRequest) (*lomsServiceApi.CreateOrderResponse, error) {
	fmt.Printf("Create order: %d items\n", len(req.Items))

	order, err := c.lomsClient.CreateOrder(ctx, req)
	if err != nil {
		return nil, err
	}

	return order, nil
}
