package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/model"
	convert "gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/repository/postgres/converter"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/repository/postgres/schema"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/repository/postgres/table"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/db"
)

func (r *repository) GetCart(ctx context.Context, userId int64) (model.Cart, error) {
	builder := sq.Select(
		"item_id", "count").
		From(table.UsersCarts).
		Where(sq.Eq{"user_id": userId}).
		PlaceholderFormat(sq.Dollar)

	var cart model.Cart

	query, args, err := builder.ToSql()
	if err != nil {
		return cart, err
	}

	q := db.Query{
		Name:     "repository.GetCart",
		QueryRaw: query,
	}

	var items []schema.Item
	if err := r.db.DB().Select(ctx, &items, q, args...); err != nil {
		return cart, err
	}

	cart.Items = convert.ToModelItems(items)

	return cart, nil
}
