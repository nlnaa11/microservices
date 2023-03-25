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

func (r *repository) RemoveFromCart(ctx context.Context, userId int64, item model.Item) error {
	// 1. Проверяем, если ли у пользователя этот товар в корзине
	count, err := r.getCountById(ctx, userId, item.Sku)
	if err != nil {
		return errors.WithMessage(err, "checking cart")
	}

	// 1.1. Если нет, легковесная ошибка
	if count == 0 {
		return libErr.ErrItemNotFound
	}

	var query string
	var args []interface{}

	if count > item.Count {
		// 2. Если после удаления, товар все еще остается, обновляем количество

		query, args, err = sq.Update(table.UsersCarts).
			Set("count", count-item.Count).
			Where(sq.Eq{"user_id": userId},
				sq.Eq{"item_id": item.Sku}).
			PlaceholderFormat(sq.Dollar).
			ToSql()
		if err != nil {
			return err
		}

	} else {
		// 3. Если после удаления, товар заканчивается, удаляем товар из корзины

		query, args, err = sq.Delete(table.UsersCarts).
			Where(sq.Eq{"user_id": userId},
				sq.Eq{"item_id": item.Sku}).
			PlaceholderFormat(sq.Dollar).
			ToSql()
		if err != nil {
			return err
		}
	}

	q := db.Query{
		Name:     "repository.RemoveFromCart",
		QueryRaw: query,
	}

	commandTag, err := r.db.DB().Exec(ctx, q, args...)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() < 1 {
		return libErr.ErrNoUpdatedOrDeletedRows
	}

	return nil
}
