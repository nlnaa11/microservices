package loms

import (
	"context"
	"fmt"
)

func (s *Service) CancelOrder(ctx context.Context, orderId uint64) error {
	fmt.Printf("the order with the number %d has been canceled\n", orderId)

	return nil
}
