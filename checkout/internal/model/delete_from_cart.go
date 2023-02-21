package model

import (
	"context"

	"github.com/pkg/errors"
)

func (m *Model) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	item := m.createItem(sku, count)
	err := m.stor.RemoveFromCart(ctx, user, *item)
	if err != nil {
		return errors.WithMessage(err, "remove from cart")
	}
	return nil
}
