package storage

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/model"
)

func (s *WrapStorage) convertToItem(ctx context.Context, modelItem model.Item) Item {
	return Item{
		Sku:   modelItem.Sku,
		Count: modelItem.Count,
	}
}

func (s *WrapStorage) convertToModelItems(ctx context.Context, items []Item) []model.Item {
	modelItems := make([]model.Item, 0, len(items))
	for _, item := range items {
		modelItems = append(modelItems, model.Item{
			Sku:   item.Sku,
			Count: item.Count,
		})
	}
	return modelItems
}

func (s *WrapStorage) convertToModelCart(ctx context.Context, cart *Cart) *model.Cart {
	modelItems := s.convertToModelItems(ctx, cart.Items)
	return &model.Cart{
		Items: modelItems,
	}
}
