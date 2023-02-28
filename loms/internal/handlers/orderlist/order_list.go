package orderlist

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

type ResponseItem struct {
	Sku   uint32 `json:"sku"`
	Count uint64 `json:"count"`
}

type Response struct {
	Status string         `json:"status"`
	User   int64          `json:"user"`
	Items  []ResponseItem `json:"items"`
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("ordedList: %+v", req)

	var response Response

	list, err := h.logic.ListOfOrder(ctx, req.OrderId)
	if err != nil {
		return response, err
	}

	response.Items = make([]ResponseItem, 0, len(list.Items))
	for _, item := range list.Items {
		response.Items = append(response.Items, ResponseItem{
			Sku:   item.Sku,
			Count: item.Count,
		})
	}
	response.Status = list.Status.String()
	response.User = list.User

	return response, nil
}
