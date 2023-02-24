package storage

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
)

type OrderData struct {
	Status uint16
	User   int64
	Items  []Item
}

func (s *WrapStorage) GetOrderData(ctx context.Context, orderId uint64) (model.OrderData, error) {
	storOrderData, err := s.orderStor.GetOrderData(ctx, orderId)
	if err != nil {
		return model.OrderData{}, errors.WithMessage(err, "getting order data")
	}

	orderData := s.convertToOrderData(ctx, storOrderData)

	return *orderData, nil
}

func (s *WrapStorage) convertToOrderData(ctx context.Context, storOrderData *OrderData) *model.OrderData {
	items := s.convertToItems(ctx, storOrderData.Items)

	return &model.OrderData{
		Status: storOrderData.Status,
		User:   storOrderData.User,
		Items:  items,
	}
}

// сюда бы дженерики
func (s *WrapStorage) convertToItems(ctx context.Context, storItems []Item) []model.Item {
	items := make([]model.Item, 0, len(storItems))

	for _, storItem := range storItems {
		items = append(items, model.Item{
			Sku:   storItem.Sku,
			Count: uint16(storItem.Count),
		})
	}

	return items
}
