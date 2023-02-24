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
	PathToCreateOrder string = "/createOrder"
)

type ItemInCart struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type CreateOrderRequest struct {
	Items []ItemInCart `json:"items"`
}

type CreateOrderResponse struct {
	OrderId int64  `json:"orderId"`
	Status  uint16 `json:"status"`
}

// TODO: validations & reserve (cancel reserve)
func (c *Client) CreateOrder(ctx context.Context, user int64, items []model.Item) (*model.Order, error) {
	var request CreateOrderRequest
	request.Items = make([]ItemInCart, 0, len(items))
	for _, item := range items {
		itemInCart := ItemInCart{
			Sku:   item.Sku,
			Count: item.Count,
		}
		request.Items = append(request.Items, itemInCart)
	}

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

	var response CreateOrderResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&response)
	if err != nil {
		return nil, errors.Wrap(err, "decoding json")
	}

	return &model.Order{
		OrderId: response.OrderId,
		Status:  response.Status,
	}, nil
}
