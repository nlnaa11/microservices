package loms

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/model"
	"gitlab.ozon.dev/nlnaa/homework-1/libs/wrappers/client"
)

const (
	pathToStocks string = "/stocks"
)

type Client struct {
	url string
}

func New(url string) *Client {
	return &Client{
		url: url,
	}
}

type StocksRequest struct {
	Sku uint32 `json:"sku"`
}

type StocksItem struct {
	WarehouseId int64  `json:"warehouseId"`
	Count       uint64 `json:"count"`
}

type StocksResponse struct {
	Stocks []StocksItem `json:"stocks"`
}

// TODO: validations
func (c *Client) Stocks(ctx context.Context, sku uint32) ([]model.Stock, error) {
	request := StocksRequest{Sku: sku}

	clientWrapper := client.New[StocksRequest, StocksResponse](c.url + pathToStocks)

	response, err := clientWrapper.Service(ctx, request)
	if err != nil {
		return nil, err
	}

	stocks := make([]model.Stock, 0, len(response.Stocks))
	for _, stock := range response.Stocks {
		stocks = append(stocks, model.Stock{
			WarehouseId: stock.WarehouseId,
			Count:       stock.Count,
		})
	}

	return stocks, nil
}
