package deletefromcart

import (
	"context"
	"log"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/service/checkout"
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
	User  int64  `json:"user"`
	Sku   uint32 `json:"sku"`
	Count uint64 `json:"count"`
}

type Response struct {
	Success bool `json:"success"`
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("deleteFromCart: %+v", req)

	var response Response

	err := h.logic.DeleteFromCart(ctx, req.User, req.Sku, req.Count)
	if err != nil {
		return response, err
	}

	return response, nil
}
