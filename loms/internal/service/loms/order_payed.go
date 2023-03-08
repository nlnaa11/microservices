package loms

import (
	"context"
	"fmt"
)

func (s *Service) OrderPayed(ctx context.Context, orderId uint64) error {
	fmt.Printf("Order #%d has been paid for\n", orderId)

	return nil
}
