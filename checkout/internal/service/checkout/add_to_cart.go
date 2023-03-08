package checkout

import (
	"context"
	"fmt"
)

func (s *Service) AddToCart(ctx context.Context, user int64, sku uint32, count uint64) error {
	fmt.Printf("Add %d of %d items to %d user cart\n", count, sku, user)

	return nil
}
