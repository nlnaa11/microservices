package loms

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	libErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/model"
)

func (s *Service) Stocks(ctx context.Context, sku uint32) ([]model.Stock, error) {
	fmt.Printf("Checking stocks to find a product with the article %d\n", sku)

	var (
		stocks        []model.Stock
		reservedCount uint64
		err           error
	)

	// 1 + 2 = transaction
	errTx := s.txManager.RepeatableRead(ctx, func(ctx context.Context) error {
		// 1. Получить запасы
		stocks, err = s.stocksRepo.CheckStocks(ctx, sku)
		if err != nil {
			return errors.WithMessage(err, "checking stocks")
		}
		// 2. Получить резервы
		reservedCount, err = s.stocksRepo.CheckReserve(ctx, sku)
		if err != nil {
			return errors.WithMessage(err, "checking reserve")
		}

		return nil
	})
	if errTx != nil {
		return nil, errTx
	}

	// 3. Найти разницу
	result, err := s.availableStocks(ctx, stocks, reservedCount)
	if err != nil {
		return nil, errors.WithMessage(err, "getting available stocks")
	}

	// !! Мы не анализируем доступность

	return result, nil
}

func (s *Service) availableStocks(ctx context.Context, stocks []model.Stock, reservedCount uint64) ([]model.Stock, error) {
	if len(stocks) == 0 {
		return nil, libErr.ErrEmptyStocks
	}

	result := make([]model.Stock, 0, len(stocks))
	var remains uint64

	for i, stock := range stocks {
		if reservedCount >= stock.Count {
			reservedCount -= stock.Count
		} else {
			remains = stock.Count - reservedCount
			reservedCount = 0
		}

		if reservedCount == 0 {
			result = stocks[i+1:]

			if remains != 0 {
				result = append(result, model.Stock{
					WarehouseId: stock.WarehouseId,
					Count:       remains,
				})
			}
			break
		}
	}

	return result, nil
}
