package storage

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
)

func (s *WrapStorage) convertToModelStocks(ctx context.Context, stocks []Stock) []model.Stock {
	modelStocks := make([]model.Stock, 0, len(stocks))

	for _, stock := range stocks {
		modelStocks = append(modelStocks, model.Stock{
			WarehouseId: stock.WarehouseId,
			Count:       stock.Count,
		})
	}

	return modelStocks
}

func (s *WrapStorage) convertToItems(ctx context.Context, modelItems []model.Item) []Item {
	items := make([]Item, 0, len(modelItems))

	for _, modelItem := range modelItems {
		items = append(items, Item{
			Sku:   modelItem.Sku,
			Count: modelItem.Count,
		})
	}

	return items
}

// сюда бы дженерики
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

func (s *WrapStorage) convertToModelOrderData(ctx context.Context, orderData *OrderData) *model.OrderInfo {
	modelItems := s.convertToModelItems(ctx, orderData.Items)

	return &model.OrderInfo{
		Order: model.Order{
			OrderId: 0,
			Status:  orderData.Status,
		},
		User:  orderData.User,
		Items: modelItems,
	}
}
