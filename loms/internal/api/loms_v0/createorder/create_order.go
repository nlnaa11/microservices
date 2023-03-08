package createorder

import (
	"context"
	"log"

	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/service/loms"
)

type Handler struct {
	logic *loms.Service
}

func New(logic *loms.Service) *Handler {
	return &Handler{
		logic: logic,
	}
}

type RequestItem struct {
	Sku   uint32 `json:"sku"`
	Count uint64 `json:"count"`
}

type Request struct {
	User  int64         `json:"user"`
	Items []RequestItem `json:"items"`
}

type Response struct {
	OrderId uint64 `json:"orderId"`
	Status  string `json:"status"`
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("createOrder: %+v", req)

	items := make([]model.Item, 0, len(req.Items))
	for _, item := range req.Items {
		items = append(items, model.Item{
			Sku:   item.Sku,
			Count: item.Count,
		})
	}

	var response Response

	orderInfo, err := h.logic.CreateOrder(ctx, req.User, items)
	if err != nil {
		return response, err
	}

	response.OrderId = orderInfo.OrderId
	response.Status = orderInfo.Status

	return response, nil
}
