package loms

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/model"
)

// OrderInfo = [orderId + status] + [userId + Items]
func (s *Service) OrderList(ctx context.Context, orderId uint64) (model.OrderInfo, error) {
	fmt.Printf("Getting an order list by id: %d", orderId)

	orderData, err := s.ordersRepo.GetOrderData(ctx, orderId)
	if err != nil {
		return model.OrderInfo{}, errors.WithMessage(err, "getting order data")
	}

	return *orderData, nil
}
