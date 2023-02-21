package addtocart

import (
	"context"
	"log"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/model"
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
	User  int64  `json:"user"`
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type Response struct {
	Success bool `json:"success"`
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("addToCart: %+v", req)

	var response Response

	err := h.logic.AddToCart(ctx, req.User, req.Sku, req.Count)
	if err != nil {
		return response, err
	}

	return response, nil
}
