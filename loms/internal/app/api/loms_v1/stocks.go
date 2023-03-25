package loms_v1

import (
	"context"

	"github.com/pkg/errors"
	desc "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/loms_v1"
	converter "gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/converter/loms_v1"
)

func (i *Implementation) Stocks(ctx context.Context, req *desc.StocksRequest) (*desc.StocksResponse, error) {
	// 1. Прошерстить склады
	stocks, err := i.lomsService.Stocks(ctx, req.GetSku())
	if err != nil {
		return nil, errors.WithMessage(err, "getting stocks")
	}

	return &desc.StocksResponse{
		Stocks: converter.ToDescStocks(stocks),
	}, nil
}
