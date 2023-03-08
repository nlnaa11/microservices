package loms

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
)

func (s *Service) Stocks(ctx context.Context, sku uint32) ([]model.Stock, error) {
	fmt.Printf("Checking stocks to find a product with the article %d\n", sku)

	stocksCount := gofakeit.IntRange(0, 64)
	stocks := make([]model.Stock, 0, stocksCount)
	for i := 0; i < stocksCount; i++ {
		stocks = append(stocks, model.Stock{
			WarehouseId: gofakeit.Int64(),
			Count:       uint64(gofakeit.IntRange(1, 100)),
		})
	}

	return stocks, nil
}
