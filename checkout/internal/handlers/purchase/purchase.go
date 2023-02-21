package purchase

import (
	"context"
	"errors"
	"log"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/model"
)

var (
	ErrInvalidOrder = errors.New("invalid order")
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
	OrderId int64
	Status  uint16
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("purchase: %+v", req)

	var response Response

	rawResponse, err := h.logic.Purchase(ctx, req.User)
	if err != nil {
		return response, err
	}
	if rawResponse == nil || rawResponse.OrderId == -1 {
		return response, ErrInvalidOrder
	}

	response.OrderId = rawResponse.OrderId
	response.Status = rawResponse.Status

	return response, nil
}
