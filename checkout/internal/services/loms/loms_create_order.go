package loms

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/model"
	"gitlab.ozon.dev/nlnaa/homework-1/libs/wrappers/client"
)

const (
	pathToCreateOrder string = "/createOrder"
)

type Item struct {
	Sku   uint32 `json:"sku"`
	Count uint64 `json:"count"`
}

type CreateOrderRequest struct {
	Items []Item `json:"items"`
}

type CreateOrderResponse struct {
	OrderId int64  `json:"orderId"`
	Status  string `json:"status"`
}

func (c *Client) CreateOrder(ctx context.Context, user int64, modelItems []model.Item) (model.OrderInfo, error) {
	var request CreateOrderRequest
	request.Items = make([]Item, 0, len(modelItems))
	for _, modelItem := range modelItems {
		request.Items = append(request.Items, Item{
			Sku:   modelItem.Sku,
			Count: uint64(modelItem.Count),
		})
	}

	clientWrapper := client.New[CreateOrderRequest, CreateOrderResponse](c.url + pathToCreateOrder)

	response, err := clientWrapper.Service(ctx, request)
	if err != nil {
		return model.OrderInfo{}, err
	}

	return model.OrderInfo{
		OrderId: response.OrderId,
		Status:  response.Status,
	}, nil
}
