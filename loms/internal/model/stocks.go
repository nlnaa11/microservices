package model

import (
	"context"

	"github.com/pkg/errors"
)

type Stock struct {
	WarehouseId int64
	Count       uint64
}

func (m *Model) Stocks(ctx context.Context, sku uint32) ([]Stock, error) {
	// резервы учтены
	stocks, err := m.stor.Stocks(ctx, sku)
	if err != nil {
		return nil, errors.WithMessage(err, "getting stocks")
	}

	return stocks, nil
}
