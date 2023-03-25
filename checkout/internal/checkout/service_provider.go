package checkout

import (
	"context"
	"log"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/repository"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/repository/postgres"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/service/checkout"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/db"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/db/transaction"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/config"
)

type serviceProvider struct {
	db db.Client

	config *config.ConfigData

	repo      repository.CartsRepository
	txManager transaction.Manager

	service *checkout.Service
}

func NewServiceProvider(config *config.ConfigData) *serviceProvider {
	return &serviceProvider{
		config: config,
	}
}

func (sp *serviceProvider) GetConfig(ctx context.Context) *config.ConfigData {
	return sp.config
}

func (sp *serviceProvider) GetDB(ctx context.Context) db.Client {
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

func (sp *serviceProvider) GetCartsRepository(ctx context.Context) *repository.CartsRepository {
	if sp.repo == nil {
		// there may be a storage selection switch here
		sp.repo = postgres.NewRepository(sp.GetDB(ctx))
	}

	return &sp.repo
}

func (sp *serviceProvider) GetTransactionManager(ctx context.Context) transaction.Manager {
	if sp.txManager == nil {
		sp.txManager = transaction.New(sp.GetDB(ctx).DB())
	}

	return sp.txManager
}

func (sp *serviceProvider) GetService(ctx context.Context) *checkout.Service {
	if sp.service == nil {
		sp.service = checkout.New(*sp.GetCartsRepository(ctx), sp.GetTransactionManager(ctx))
	}

	return sp.service
}
