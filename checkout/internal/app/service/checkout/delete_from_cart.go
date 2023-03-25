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
func (s *Service) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint64) error {
	fmt.Printf("Delete %d of %d items from %d user\n", count, sku, user)

	errTx := s.txManager.Serializable(ctx, func(ctx context.Context) error {
		err := s.cartsRepo.RemoveFromCart(ctx, user, model.Item{Sku: sku, Count: count})
		if err != nil {
			return errors.WithMessage(err, "deleting item from cart")
		}
		return nil
	})
	if errTx != nil {
		return errTx
	}

	return nil
}
