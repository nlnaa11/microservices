package converter

import (
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/model"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/repository/postgres/schema"
)

func ToModelItems(items []schema.Item) []model.Item {
	result := make([]model.Item, 0, len(items))
	for _, item := range items {
		result = append(result, model.Item{
			Sku:   item.Id,
			Count: item.Count,
		})
	}

	return result
}
