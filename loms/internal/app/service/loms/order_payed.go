package loms

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/model"
)

// TODO: подумать, как лучше, как добиться 100 % удаления

func (s *Service) OrderPayed(ctx context.Context, orderId uint64) error {
	fmt.Printf("Order #%d has been paid for\n", orderId)

	// 1. Установить статус
	_ = s.ordersRepo.SetOrderStatus(ctx, orderId, model.StatusPayed.String())

	// 2. Получить элементы заказа
	items, err := s.ordersRepo.GetOrderItems(ctx, orderId)
	if err != nil {
		return err
	}

	// 3. Подготовить данные для удаления запасов
	// 4. Удалить из запасов
	// 5. Удалить из резерва
	// [3:5] = transaction
	for _, item := range items {
		errTx := s.txManager.RepeatableRead(ctx, func(ctx context.Context) error {
			whIndicesToRemove, whToUpdate, err := s.prepareData(ctx, item)
			if err != nil {
				// todo: handle this problem normally
				return err
			}
			err = s.stocksRepo.RemoveFromStocks(ctx, item.Sku, whIndicesToRemove, whToUpdate)
			if err != nil {
				return err
			}

			err = s.stocksRepo.RemoveFromReserve(ctx, []model.Item{item})
			if err != nil {
				return err
			}

			return nil
		})
		if errTx != nil {
			return errTx
		}
	}

	return nil
}

// слайс индексов складов на удаление + склад (индекс и количество) на обновление
func (s *Service) prepareData(ctx context.Context, item model.Item) ([]int64, *model.Stock, error) {
	stocks, err := s.stocksRepo.CheckStocks(ctx, item.Sku)
	if err != nil {
		return nil, nil, errors.WithMessage(err, "checking stocks")
	}

	var whIndicesToRemove []int64
	var whToUpdate model.Stock

	remains := item.Count

	for _, stock := range stocks {
		if stock.Count <= remains {
			whIndicesToRemove = append(whIndicesToRemove, stock.WarehouseId)
			remains -= stock.Count
		} else {
			whToUpdate.WarehouseId = stock.WarehouseId
			whToUpdate.Count = stock.Count - remains
		}

		if remains == 0 {
			break
		}
	}

	return whIndicesToRemove, &whToUpdate, nil
}
