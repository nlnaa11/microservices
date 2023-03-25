package purchase

import (
	"context"
	"log"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/service/checkout"
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

type Response struct {
	OrderId int64  `json:"orderId"`
	Status  string `json:"status"`
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("purchase: %+v", req)

	var response Response

	_, err := h.logic.Purchase(ctx, req.User)
	if err != nil {
		return response, err
	}

	return response, nil
}
