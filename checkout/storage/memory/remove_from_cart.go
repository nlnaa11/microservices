package memory

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/storage"
	"gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

func (m *MemoryWrapper) RemoveFromCart(ctx context.Context, storItem storage.Item, user int64) error {
	if m.carts == nil {
		return errors.ErrNoBusyCart
	}

	cart, ok := m.carts[user]
	if ok {
		for i, item := range cart.items {
			if item.sku == storItem.Sku {
				if item.count <= storItem.Count {
					cart.items[i].count = 0
					break
				}
				cart.items[i].count -= storItem.Count
				break
			}
		}
	}
	return nil
}
