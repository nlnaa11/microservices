package postgres

import (
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/clients/db"
)

type repository struct {
	db db.Client
}

func NewRepository(db db.Client) *repository {
	return &repository{
		db: db,
	}
}
