package memory

import (
	"context"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/storage"
	"gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

func (m *MemoryWrapper) GetAll(ctx context.Context, user int64) (*storage.Cart, error) {
	if m.carts == nil {
		return nil, errors.ErrNoBusyCart
	}

	if cart, ok := m.carts[user]; ok {
		return m.convertToStorCart(ctx, cart), nil
	}

	return nil, errors.ErrEmptyCart
}
