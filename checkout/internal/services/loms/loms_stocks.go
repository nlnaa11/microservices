package loms

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/model"
)

const (
	PathStocks string = "/stocks"
)

type Client struct {
	url     string
	urlGoal string
}

func New(url string, pathTo string) *Client {
	return &Client{
		url:     url,
		urlGoal: url + pathTo,
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

	rawJSON, err := json.Marshal(request)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling json")
	}

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, c.urlGoal, bytes.NewBuffer(rawJSON))
	if err != nil {
		return nil, errors.Wrap(err, "creating http request")
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return nil, errors.Wrap(err, "calling http")
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", httpResponse.StatusCode)
	}

	var response StocksResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&response)
	if err != nil {
		return nil, errors.Wrap(err, "decoding json")
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
