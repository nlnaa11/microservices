package purchase

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

type Response struct {
	OrderId int64  `json:"orderId"`
	Status  string `json:"status"`
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("purchase: %+v", req)

	var response Response

	rawResponse, err := h.logic.Purchase(ctx, req.User)
	if err != nil {
		return response, err
	}
	if rawResponse.Status == "unknown" || rawResponse.OrderId == 0 {
		return response, errors.ErrInvalidOrder
	}

	response.OrderId = rawResponse.OrderId
	response.Status = rawResponse.Status

	return response, nil
}
