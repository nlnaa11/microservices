package product

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/model"
	"gitlab.ozon.dev/nlnaa/homework-1/libs/wrappers/client"
)

const (
	pathToProducts string = "/docs/"
)

type Client struct {
	url string
}

func New(url string) *Client {
	return &Client{
		url: url,
	}
}

type ProductRequest struct {
	Token string `json:"token"`
	Sku   uint32 `json:"sku"`
}

// подумать о запросе информации сразу о нескольких товарах
type ProductResponse struct {
	Sku   uint32  `json:"sku"`
	Count uint64  `json:"count"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// validations
func (c *Client) GetProductInfo(ctx context.Context, token string, sku uint32) (model.ItemInfo, error) {
	request := ProductRequest{
		Token: token,
		Sku:   sku,
	}

	clientWrapper := client.New[ProductRequest, ProductResponse](c.url + pathToProducts)

	response, err := clientWrapper.Service(ctx, request)
	if err != nil {
		return model.ItemInfo{}, err
	}

	return model.ItemInfo{
		Item: model.Item{
			Sku:   response.Sku,
			Count: response.Count,
		},
		Name:  response.Name,
		Price: response.Price,
	}, nil
}
