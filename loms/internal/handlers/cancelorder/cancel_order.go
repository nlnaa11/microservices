package cancelorder

import (
	"context"
	"log"

	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
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
	OrderId uint64 `json:"orderId"`
}

type Response struct {
	Success bool `json:"success"`
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("cancelOrder: %+v", req)

	var response Response

	err := h.logic.CancelOrder(ctx, req.OrderId)
	if err != nil {
		return response, err
	}

	response.Success = true

	return response, nil
}
