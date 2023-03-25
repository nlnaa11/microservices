package orderpayed

import (
	"context"
	"log"

	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/service/loms"
)

type Handler struct {
	logic *loms.Service
}

func New(logic *loms.Service) *Handler {
	return &Handler{
		logic: logic,
	}
}

type Request struct {
	OrderId uint64 `json:"orderId"`
}

type Response struct {
	Success bool `json:"success"`
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("orderPayed: %+v", req)

	var response Response

	err := h.logic.OrderPayed(ctx, req.OrderId)
	if err != nil {
		return response, err
	}

	response.Success = true
	return response, nil
}
