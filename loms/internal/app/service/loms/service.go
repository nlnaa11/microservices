package loms

import (
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/repository"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/clients/db/transaction"
)

type Service struct {
	ordersRepo repository.OrdersRepository
	stocksRepo repository.StocksRepository

	txManager transaction.Manager
}

func New(ordersRepo repository.OrdersRepository, stockRepo repository.StocksRepository, txManager transaction.Manager) *Service {
	return &Service{
		ordersRepo: ordersRepo,
		stocksRepo: stockRepo,
		txManager:  txManager,
	}
}
