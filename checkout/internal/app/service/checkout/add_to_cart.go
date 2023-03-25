package checkout

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/model"
)

// Пользователь может менять решение о количестве товара быстро и часто.
// Это наводит на мысль об обертке всей операции в Транзакцию.
// Добавление + удаление = Serializable (TxIsoLevel)
func (s *Service) AddToCart(ctx context.Context, user int64, sku uint32, count uint64) error {
	fmt.Printf("Add %d of %d items to %d user cart\n", count, sku, user)

	errTx := s.txManager.Serializable(ctx, func(ctx context.Context) error {
		err := s.cartsRepo.AddToCart(ctx, user, model.Item{Sku: sku, Count: count})
		if err != nil {
			return errors.WithMessage(err, "adding item to cart")
		}
		return nil
	})
	if errTx != nil {
		return errTx
	}

	return nil
}
