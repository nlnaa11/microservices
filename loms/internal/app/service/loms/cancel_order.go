package loms

import (
	"context"
	"fmt"

	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/model"
)

// TODO: обеспечить 99 %-ое удаление

func (s *Service) CancelOrder(ctx context.Context, orderId uint64) error {
	fmt.Printf("the order with the number %d has been canceled\n", orderId)

	// 1. Установить статус
	_ = s.ordersRepo.SetOrderStatus(ctx, orderId, model.StatusCancelled.String())

	// 2. Получить товары заказа
	items, err := s.ordersRepo.GetOrderItems(ctx, orderId)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}

	// 3. Удалить товары из резерва
	// (разделила с расчетом на распараллеливание. Возможно, снова объединю)
	for _, item := range items {
		err = s.removeFromReserve(ctx, &item)
		if err != nil {
			return err
		}
	}

	return nil
}

// мне кажется, так будет проще распараллелить выполнениe задачи
func (s *Service) removeFromReserve(ctx context.Context, item *model.Item) error {
	errTx := s.txManager.RepeatableRead(ctx, func(ctx context.Context) error {
		err := s.stocksRepo.RemoveFromReserve(ctx, []model.Item{*item})
		if err != nil {
			return err
		}
		return nil
	})
	if errTx != nil {
		return errTx
	}

	return nil
}
