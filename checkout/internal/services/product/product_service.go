package product

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
	PathProducts string = "/docs/"
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

type ProductRequest struct {
	Sku uint32 `json:"sku"`
}

// подумать о запросе информации сразу о нескольких товарах
type ProductResponse struct {
	Sku   uint32  `json:"sku"`
	Count uint16  `json:"count"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// validations
func (c *Client) GetProductInfo(ctx context.Context, sku uint32) (model.ItemInfo, error) {
	request := ProductRequest{Sku: sku}

	var productInfo model.ItemInfo

	rawJSON, err := json.Marshal(request)
	if err != nil {
		return productInfo, errors.Wrap(err, "marshaling json")
	}

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, c.urlGoal, bytes.NewBuffer(rawJSON))
	if err != nil {
		return productInfo, errors.Wrap(err, "creating http request")
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return productInfo, errors.Wrap(err, "calling http")
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return productInfo, fmt.Errorf("wrong status code: %d", httpResponse.StatusCode)
	}

	var response ProductResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&response)
	if err != nil {
		return productInfo, errors.Wrap(err, "decoding json")
	}

	productInfo.Sku = response.Sku
	productInfo.Count = response.Count
	productInfo.Name = response.Name
	productInfo.Price = response.Price

	return productInfo, nil
}
