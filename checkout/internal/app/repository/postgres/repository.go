package postgres

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/repository/postgres/table"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/db"
)

type repository struct {
	db db.Client
}

func NewRepository(db db.Client) *repository {
	return &repository{
		db: db,
	}
}

// в корзине не может быть 0 товаров какого-л типа.
// либо товар есть, либо его нет
func (r *repository) getCountById(ctx context.Context, userId int64, itemId uint32) (uint64, error) {
	builder := sq.Select("count").
		From(table.UsersCarts).
		Where(sq.Eq{"user_id": userId}).
		Where(sq.Eq{"item_id": itemId}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "repository.getCountById",
		QueryRaw: query,
	}

	var count uint64
	if err := r.db.DB().Select(ctx, &count, q, args...); err != nil {
		return 0, err
	}

	log.Printf("found %d items with sku %d in the cart\n", count, itemId)

	return count, nil
}
