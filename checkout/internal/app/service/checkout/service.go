package checkout

import (
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/repository"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/db/transaction"
)

type Service struct {
	cartsRepo repository.CartsRepository

	// todo: подумать, как пожно сделать по-другому
	txManager transaction.Manager
}

func New(cartsRepo repository.CartsRepository, txManager transaction.Manager) *Service {
	return &Service{
		cartsRepo: cartsRepo,
		txManager: txManager,
	}
}
