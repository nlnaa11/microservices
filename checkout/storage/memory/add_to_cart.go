package memory

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/storage"
	"gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

func (m *MemoryWrapper) AddToCart(ctx context.Context, storItem storage.Item, user int64) error {
	if m.carts == nil {
		return errors.ErrNoBusyCart
	}

	item := m.convertToItem(ctx, &storItem)

	cart, ok := m.carts[user]
	if ok {
		cart.items = append(cart.items, *item)
		return nil
	}

	items := make([]Item, 0, 1)
	items = append(items, *item)
	m.carts[user].items = items

	return nil
}
