package loms

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
	libErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/model"
)

func (s *Service) CreateOrder(ctx context.Context, user int64, items []model.Item) (model.Order, error) {
	fmt.Printf("Creating order fo %d user from %d products\n", user, len(items))

	order := model.Order{
		OrderId: 0,
		Status:  model.StatusFailed.String(),
	}

	if len(items) == 0 {
		return order, libErr.ErrEmptyItems
	}

	// 3. Cоздать заказ: регистрация нового + сохранение в истории
	orderId, err := s.ordersRepo.CreateOrder(ctx, user, items)
	if errors.As(err, libErr.ErrOrderRegistation) {
		return order, errors.WithMessage(err, "creating order")
	}
	if err != nil {
		_ = s.ordersRepo.SetOrderStatus(ctx, orderId, model.StatusFailed.String())
		return order, errors.WithMessage(err, "creating order")
	}

	// order -- не какая-то постоянная сущность, поэтому можем передать только id
	// для отслеживания статуса
	// мы не будем ждать завершения. Обрабатывать ошибку нужно др способом
	go s.cancelByTimeout(ctx, orderId)

	for _, item := range items {
		// 1. Проверить запасы: стоки - резерв
		// 2. Оценить запасы: [резервация и ожидание оплаты] || отмена оплаты
		errTx := s.txManager.RepeatableRead(ctx, func(ctx context.Context) error {
			_, err = s.estimateStocks(ctx, []model.Item{item})
			if err != nil {
				return errors.WithMessage(err, "checking stocks")
			}

			err = s.stocksRepo.AddToReserve(ctx, []model.Item{item})
			if err != nil {
				return errors.WithMessage(err, "reserving items")
			}

			return nil
		})
		if errTx != nil {
			_ = s.ordersRepo.SetOrderStatus(ctx, orderId, model.StatusFailed.String())
			return order, errTx
		}

	}

	_ = s.ordersRepo.SetOrderStatus(ctx, orderId, model.StatusAwaitingPayment.String())

	return order, nil
}

// если честно, создавать для каждого заказа свою горутину и свой таймер --
// по-моему, это чересчур
// зато просто)
func (s *Service) cancelByTimeout(ctx context.Context, orderId uint64) {
	defer func() {
		fmt.Println("cancelByTimeout completed")
	}()

	timer := time.NewTimer(10 * time.Minute)

	select {
	case <-ctx.Done():
		log.Println("something was wrong")
		return
	case <-timer.C:
		order, _ := s.ordersRepo.GetOrderInfo(ctx, orderId)
		if order.Status == model.StatusAwaitingPayment.String() {
			_ = s.CancelOrder(ctx, order.OrderId)
		}
	}
}

// NB: часть уже былa написана, зачем добру пропадать
// возвращает ошибку, если хотя бы по одному товару недостаток.
// при ошибке также возвращает возможное количество тех товаров,
// которых недостаточно для заказа.
// если всех товаров достаточно, оба возвращаемых значения nil
func (s *Service) estimateStocks(ctx context.Context, items []model.Item) (map[uint32]uint64, error) {
	if len(items) == 0 {
		return nil, libErr.ErrEmptyItems
	}

	inDeficit := make(map[uint32]uint64)

	for _, item := range items {
		stocks, err := s.Stocks(ctx, item.Sku)
		if err != nil {
			inDeficit[item.Sku] = 0

			log.Printf("%s\n", err.Error())
			continue
		}

		var availabilityCount int64
		requiredCount := int64(item.Count)
		isEnough := false

		for _, stock := range stocks {
			requiredCount -= int64(stock.Count)
			if requiredCount <= 0 {
				isEnough = true
				break
			}
			availabilityCount += int64(stock.Count)
		}

		if !isEnough {
			inDeficit[item.Sku] = uint64(availabilityCount)
		}
	}

	if len(inDeficit) != 0 {
		return inDeficit, errors.New("deficits of items")
	}

	return nil, nil
}
