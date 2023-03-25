package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/model"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/app/repository/postgres/table"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/db"
	libErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

func (r *repository) AddToCart(ctx context.Context, userId int64, item model.Item) error {
	// 1. Проверяем, если ли уже у пользователя этот товар в корзине
	count, err := r.getCountById(ctx, userId, item.Sku)
	if err != nil {
		return errors.WithMessage(err, "checking cart")
	}

	var query string
	var args []interface{}

	if count == 0 {
		// 2. Если нет, добавляем товар в корзину

		query, args, err = sq.Insert(table.UsersCarts).
			Columns("user_id", "item_id", "count").
			Values(userId, item.Sku, item.Count).
			PlaceholderFormat(sq.Dollar).
			ToSql()
		if err != nil {
			return err
		}

	} else {
		// 3. Если есть, обновляем количество

		query, args, err = sq.Update(table.UsersCarts).
			Set("count", count+item.Count).
			Where(sq.Eq{"user_id": userId},
				sq.Eq{"item_id": item.Sku}).
			PlaceholderFormat(sq.Dollar).
			ToSql()
		if err != nil {
			return err
		}
	}

	q := db.Query{
		Name:     "repository.AddToCart",
		QueryRaw: query,
	}

	commandTag, err := r.db.DB().Exec(ctx, q, args...)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() < 1 {
		return libErr.ErrNoUpdatedOrAddedRows
	}

	return nil
}
