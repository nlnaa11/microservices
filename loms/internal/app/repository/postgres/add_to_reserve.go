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

func (r *repository) AddToReserve(ctx context.Context, items []model.Item) error {
	if len(items) == 0 {
		return libErr.ErrEmptyItems
	}

	for _, item := range items {
		err := r.addToReserve(ctx, item)
		if err != nil {
			return errors.WithMessage(err, "adding to reserve")
		}
	}

	return nil
}

func (r *repository) addToReserve(ctx context.Context, item model.Item) error {
	// 1. Проверяем, зарезервирован ли товар
	count, err := r.getCountById(ctx, item.Sku)
	if err != nil {
		return errors.WithMessage(err, "checking reserve")
	}

	var query string
	var args []interface{}

	if count == 0 {
		// 2. Если нет, добавляем товар в корзину

		query, args, err = sq.Insert(table.ReservedItems).
			Columns("item_id", "count").
			Values(item.Sku, item.Count).
			PlaceholderFormat(sq.Dollar).
			ToSql()
		if err != nil {
			return err
		}

	} else {
		// 3. Если есть, обновляем количество

		query, args, err = sq.Update(table.ReservedItems).
			Set("count", count+item.Count).
			Where(sq.Eq{"item_id": item.Sku}).
			PlaceholderFormat(sq.Dollar).
			ToSql()
		if err != nil {
			return err
		}
	}

	q := db.Query{
		Name:     "repository.addToReserve",
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

// в резерве не может быть 0 товаров какого-л типа.
// либо товар есть, либо его нет
func (r *repository) getCountById(ctx context.Context, itemId uint32) (uint64, error) {
	builder := sq.Select("count").
		From(table.ReservedItems).
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

	return count, nil
}
