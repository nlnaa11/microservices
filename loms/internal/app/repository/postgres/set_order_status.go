package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	libErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/repository/postgres/schema"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/repository/postgres/table"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/clients/db"
)

func (r *repository) SetOrderStatus(ctx context.Context, orderId uint64, status string) error {
	builder := sq.Update(table.Orders).
		Set("status", schema.StatusFromString(status)).
		Where(sq.Eq{"order_id": orderId}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "repository.SetOrderStatus",
		QueryRaw: query,
	}

	commandTag, err := r.db.DB().Exec(ctx, q, args...)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() < 1 {
		return libErr.ErrNoUpdatedRows
	}

	return nil
}
