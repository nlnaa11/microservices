package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/repository/postgres/table"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/clients/db"
)

// invalid count = 0
func (r *repository) CheckReserve(ctx context.Context, itemId uint32) (uint64, error) {
	builder := sq.Select("count").
		From(table.ReservedItems).
		Where(sq.Eq{"item_id": itemId}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "repository.CheckReserve",
		QueryRaw: query,
	}

	var count uint64
	if err := r.db.DB().Select(ctx, &count, q, args...); err != nil {
		return 0, err
	}

	return count, nil
}
