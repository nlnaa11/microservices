package storage

import (
	"context"
	"log"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/model"
)

func (s *WrapStorage) RemoveFromCart(ctx context.Context, user int64, item model.Item) error {
	ctx, cancel := context.WithCancel(ctx)

	defer func() {
		log.Println("RemoveFromCart complete")
		cancel()
	}()

	storItem := *s.createStorItem(item)
	successChan := make(chan bool)
	go s.cartStor.RemoveFrom(storItem, user, successChan)

	select {
	case <-ctx.Done():
		// TODO
		return ErrCancel
	case success := <-successChan:
		if !success {
			return ErrAccessToStor
		}
		return nil
	}
}
