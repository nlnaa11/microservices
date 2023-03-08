package loms_v1

import (
	desc "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/loms_v1"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
)

func ToDescStocks(stocks []model.Stock) []*desc.StockItem {
	res := make([]*desc.StockItem, 0, len(stocks))
	for _, stock := range stocks {
		res = append(res, ToDescStock(&stock))
	}

	return res
}

func ToDescStock(stock *model.Stock) *desc.StockItem {
	return &desc.StockItem{
		WarehouseId: stock.WarehouseId,
		Count:       stock.Count,
	}
}
