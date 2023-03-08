package loms

import (
	"context"
	"fmt"

	lomsServiceApi "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/loms_v1"
)

func (c *client) Stocks(ctx context.Context, req *lomsServiceApi.StocksRequest) (*lomsServiceApi.StocksResponse, error) {
	fmt.Printf("Stocks: inventory of %d products\n", req.Sku)

	stocks, err := c.lomsClient.Stocks(ctx, req)
	if err != nil {
		return nil, err
	}

	return stocks, nil
}
