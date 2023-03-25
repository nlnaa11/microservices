package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/repository/postgres/table"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/db"
	libErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

func (r *repository) ClearCart(ctx context.Context, userId int64) error {
	builder := sq.Delete(table.UsersCarts).
		Where(sq.Eq{"user_id": userId}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "repository.CleaCart",
		QueryRaw: query,
	}

	commandTag, err := r.db.DB().Exec(ctx, q, args...)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() < 1 {
		return libErr.ErrNoDeletedRows
	}

	return nil
}
