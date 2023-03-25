package converter

import (
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/model"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/repository/postgres/schema"
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

func ToSchemaItems(items []model.Item) []schema.Item {
	result := make([]schema.Item, 0, len(items))
	for _, item := range items {
		result = append(result, schema.Item{
			Id:    item.Sku,
			Count: item.Count,
		})
	}

	return result
}
