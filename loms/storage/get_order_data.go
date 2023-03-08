package storage

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
)

type OrderData struct {
	Status string
	User   int64
	Items  []Item
}

// возвращает данные заказа, включая его статус, в случае успеха
func (s *WrapStorage) GetOrderData(ctx context.Context, orderId uint64) (model.OrderInfo, error) {
	orderData, err := s.orderStor.GetOrderData(ctx, orderId)
	if err != nil {
		return model.OrderInfo{}, errors.WithMessage(err, "getting order data")
	}

	modelOrderData := s.convertToModelOrderData(ctx, orderData)

	return *modelOrderData, nil
}
