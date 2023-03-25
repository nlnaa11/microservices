package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	libErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/model"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/repository/postgres/table"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/clients/db"
)

func (r *repository) RemoveFromReserve(ctx context.Context, items []model.Item) error {
	if len(items) == 0 {
		return libErr.ErrEmptyItems
	}

	for _, item := range items {
		err := r.removeFromReserve(ctx, item)
		if err != nil {
			return errors.WithMessage(err, "removing from reserve")
		}
	}

	return nil
}

func (r *repository) removeFromReserve(ctx context.Context, item model.Item) error {
	// 1. Проверяем, зарезервирован ли товар
	count, err := r.getCountById(ctx, item.Sku)
	if err != nil {
		return errors.WithMessage(err, "checking reserve")
	}

	// 1.1. Если нет -- легковесная ошибка
	if count == 0 {
		return libErr.ErrItemNotFound
	}

	var query string
	var args []interface{}

	if count > item.Count {
		// 2. Если после удаления, товар все еще остается, обновляем количество

		query, args, err = sq.Update(table.ReservedItems).
			Set("count", count-item.Count).
			Where(sq.Eq{"item_id": item.Sku}).
			PlaceholderFormat(sq.Dollar).
			ToSql()
		if err != nil {
			return err
		}

	} else {
		// 3. Если после удаления, товар заканчивается, удаляем товар из корзины

		query, args, err = sq.Delete(table.ReservedItems).
			Where(sq.Eq{"item_id": item.Sku}).
			PlaceholderFormat(sq.Dollar).
			ToSql()
		if err != nil {
			return err
		}
	}

	q := db.Query{
		Name:     "repository.removeFromReserve",
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
