package checkout

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/model"
	internalErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

func (s *Service) CartList(ctx context.Context, user int64) (model.CartList, error) {
	fmt.Printf("A cart list of the %d user\n", user)

	var cartList model.CartList

	cart, err := s.cartsRepo.GetCart(ctx, user)
	if err != nil {
		return cartList, errors.WithMessage(err, "getting cart")
	}
	if len(cart.Items) == 0 {
		return cartList, internalErr.ErrEmptyCart
	}

	cartList.ItemsInfo = make([]model.ItemInfo, 0, len(cart.Items))
	for _, item := range cart.Items {
		cartList.ItemsInfo = append(cartList.ItemsInfo, model.ItemInfo{
			Item: item,
		})
	}

	return cartList, nil
}
