package checkout

import (
	"context"
	"fmt"
)

func (s *Service) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint64) error {
	fmt.Printf("Delete %d of %d items from %d user\n", count, sku, user)

	return nil
}
