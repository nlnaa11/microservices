package storage

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/model"
)

var (
	ErrAccessToStor = errors.New("memory access problems")
	ErrEmpryCart    = errors.New("empty cart")
	ErrCancel       = errors.New("operations is canceled")
)

type Item struct {
	Sku   int32
	Count int16
}

func (s *WrapStorage) AddToCart(ctx context.Context, user int64, item model.Item) error {
	ctx, cancel := context.WithCancel(ctx)

	defer func() {
		log.Println("AddToCart complete")
		cancel()
	}()

	storItem := *s.createStorItem(item)
	successChan := make(chan bool, 1)
	go s.cartStor.AddTo(storItem, user, successChan)

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

func (s *WrapStorage) createStorItem(item model.Item) *Item {
	return &Item{
		Sku:   int32(item.Sku),
		Count: int16(item.Count),
	}
}
