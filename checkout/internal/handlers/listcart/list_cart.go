package listcart

import (
	"context"
	"log"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/model"
	"gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

type Handler struct {
	logic *model.Model
}

func New(logic *model.Model) *Handler {
	return &Handler{
		logic: logic,
	}
}

type Request struct {
	User int64 `json:"user"`
}

type ItemInfo struct {
	Sku   uint32  `json:"sku"`
	Count uint64  `json:"count"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Response struct {
	Items      []ItemInfo `json:"items"`
	TotalPrice float64    `json:"totalPrice"`
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("listCart: %+v", req)

	var response Response

	rawResponse, err := h.logic.ListCart(ctx, req.User)
	if err != nil {
		return response, err
	}
	if rawResponse == nil {
		return response, errors.ErrEmptyListOfCart
	}

	response.Items = make([]ItemInfo, 0, len(rawResponse.ItemsInfo))
	for _, itemInfo := range rawResponse.ItemsInfo {
		response.Items = append(response.Items, ItemInfo{
			Sku:   itemInfo.Sku,
			Count: itemInfo.Count,
			Name:  itemInfo.Name,
			Price: itemInfo.Price,
		})
	}
	response.TotalPrice = rawResponse.TotalPrice

	return response, nil
}
