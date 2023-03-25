package converter

import (
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/model"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/repository/postgres/schema"
)

func ToModelStocks(stocks []schema.Stock) []model.Stock {
	result := make([]model.Stock, 0, len(stocks))
	for _, stock := range stocks {
		result = append(result, model.Stock{
			WarehouseId: stock.WarehouseId,
			Count:       stock.Count,
		})
	}

	return result
}
