package model

import (
	"context"

	"github.com/pkg/errors"
)

func (m *Model) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint64) error {
	item := Item{
		Sku:   sku,
		Count: count,
	}
	if err := m.stor.RemoveFromCart(ctx, user, item); err != nil {
		return errors.WithMessage(err, "remove from cart")
	}

	return nil
}
