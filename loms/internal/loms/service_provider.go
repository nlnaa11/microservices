package loms

import (
	"context"
	"log"

	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/repository"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/repository/postgres"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/service/loms"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/clients/db"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/clients/db/transaction"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/config"
)

type serviceProvider struct {
	db db.Client

	config *config.ConfigData

	ordersRepo repository.OrdersRepository
	stocksRepo repository.StocksRepository
	txManager  transaction.Manager

	service *loms.Service
}

func NewServiceProvider(config *config.ConfigData) *serviceProvider {
	return &serviceProvider{
		config: config,
	}
}

func (sp *serviceProvider) GetConfig(ctx context.Context) *config.ConfigData {
	return sp.config
}

func (sp *serviceProvider) GetDD(ctx context.Context) db.Client {
	if sp.db == nil {
		dbConfig, err := sp.GetConfig(ctx).GetDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		dbClient, err := db.New(ctx, dbConfig)
		if err != nil {
			log.Fatalf("failed to connect to db: %s", err.Error())
		}

		sp.db = dbClient
	}

	return sp.db
}

func (sp *serviceProvider) GetOrdersRepository(ctx context.Context) repository.OrdersRepository {
	if sp.ordersRepo == nil {
		sp.ordersRepo = postgres.NewRepository(sp.GetDD(ctx))
	}

	return sp.ordersRepo
}

func (sp *serviceProvider) GetStocksRepository(ctx context.Context) repository.StocksRepository {
	if sp.stocksRepo == nil {
		sp.stocksRepo = postgres.NewRepository(sp.GetDD(ctx))
	}

	return sp.stocksRepo
}

func (sp *serviceProvider) GetTransactionManager(ctx context.Context) transaction.Manager {
	if sp.txManager == nil {
		sp.txManager = transaction.New(sp.GetDD(ctx).DB())
	}

	return sp.txManager
}

func (sp *serviceProvider) GetService(ctx context.Context) *loms.Service {
	if sp.service == nil {
		sp.service = loms.New(
			sp.GetOrdersRepository(ctx),
			sp.GetStocksRepository(ctx),
			sp.GetTransactionManager(ctx))
	}

	return sp.service
}
