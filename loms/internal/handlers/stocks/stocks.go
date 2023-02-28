package stocks

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
	Sku uint32 `json:"sku"`
}

type ResponseItem struct {
	WarehouseId int64  `json:"warehouseId"`
	Count       uint64 `json:"count"`
}

type Response struct {
	Stocks []ResponseItem `json:"stocks"`
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("stocks: %+v", req)

	var response Response

	stocks, err := h.logic.Stocks(ctx, req.Sku)
	if err != nil {
		return response, err
	}

	response.Stocks = make([]ResponseItem, 0, len(stocks))
	for _, stock := range stocks {
		response.Stocks = append(response.Stocks, ResponseItem{
			WarehouseId: stock.WarehouseId,
			Count:       stock.Count,
		})
	}

	return response, nil
}
