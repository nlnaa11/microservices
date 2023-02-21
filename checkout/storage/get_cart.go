package storage

import (
	"context"
	"log"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/model"
)

type Cart struct {
	Items []Item
}

func (s *WrapStorage) GetCart(ctx context.Context, user int64) (model.Cart, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		log.Println("GetCart complete")
		cancel()
	}()

	cartChan := make(chan Cart)
	go func(cartChan chan Cart) {
		cart, err := s.cartStor.Get(user)
		if err != nil {
			log.Println(err)
		}

		if cartChan != nil {
			if cart != nil {
				cartChan <- *cart
				return
			}
			cartChan <- Cart{}
		}
	}(cartChan)

	select {
	case <-ctx.Done():
		// TODO
		return model.Cart{}, ErrCancel
	case cart := <-cartChan:
		return *s.createCart(&cart), nil
	}
}

func (s *WrapStorage) createCart(storCart *Cart) *model.Cart {
	var cart model.Cart
	cart.Items = make([]model.Item, 0, len(cart.Items))
	for _, storItem := range cart.Items {
		item := model.Item{
			Sku:   uint32(storItem.Sku),
			Count: uint16(storItem.Count),
		}
		cart.Items = append(cart.Items, item)
	}

	return &cart
}
