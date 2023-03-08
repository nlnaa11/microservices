package loms

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
)

func (s *Service) CreateOrder(ctx context.Context, user int64, items []model.Item) (model.Order, error) {
	fmt.Printf("Creating order fo %d user from %d products\n", user, len(items))

	return model.Order{
		OrderId: gofakeit.Uint64(),
		Status:  model.StatusNew.String(),
	}, nil

}
