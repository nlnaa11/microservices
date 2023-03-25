package cartlist

import (
	"context"
	"log"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/service/checkout"
	"gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

type Handler struct {
	logic *checkout.Service
}

func New(logic *checkout.Service) *Handler {
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

	rawResponse, err := h.logic.CartList(ctx, req.User)
	if err != nil {
		return response, err
	}
	if len(rawResponse.ItemsInfo) == 0 {
		return response, errors.ErrEmptyListOfCart
	}

	response.Items = make([]ItemInfo, 0, len(rawResponse.ItemsInfo))
	for _, itemInfo := range rawResponse.ItemsInfo {
		response.Items = append(response.Items, ItemInfo{
			Sku:   itemInfo.Item.Sku,
			Count: itemInfo.Item.Count,
			Name:  itemInfo.Name,
			Price: itemInfo.Price,
		})
	}
	response.TotalPrice = rawResponse.TotalPrice

	return response, nil
}
