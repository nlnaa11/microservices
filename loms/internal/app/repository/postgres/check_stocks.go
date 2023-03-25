package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/model"
	convert "gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/repository/postgres/converter"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/repository/postgres/schema"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/repository/postgres/table"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/clients/db"
)

func (r *repository) CheckStocks(ctx context.Context, itemId uint32) ([]model.Stock, error) {
	builder := sq.Select(
		"warehouse_id", "count").
		From(table.ItemsStocks).
		Where(sq.Eq{"item_id": itemId}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "repository.CheckStocks",
		QueryRaw: query,
	}

	var stocks []schema.Stock
	if err := r.db.DB().Select(ctx, &stocks, q, args...); err != nil {
		return nil, err
	}

	result := convert.ToModelStocks(stocks)

	return result, nil
}
