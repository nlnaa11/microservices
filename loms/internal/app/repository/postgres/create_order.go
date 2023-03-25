package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	libErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/model"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/repository/postgres/schema"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/repository/postgres/table"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/clients/db"
)

// invalid order: orderId = 0
func (r *repository) CreateOrder(ctx context.Context, userId int64, items []model.Item) (uint64, error) {
	// 1. Зарегистрировать новый заказ
	orderId, err := r.registerOrder(ctx)
	if err != nil {
		return 0, errors.WithMessage(err, "order registration")
	}
	// 2. Сохранить в истории
	err = r.saveOrderItems(ctx, orderId, items)
	if err != nil {
		return 0, errors.WithMessage(err, "saving order items")
	}
	err = r.saveOrderUser(ctx, orderId, userId)
	if err != nil {
		return 0, errors.WithMessage(err, "saving order user")
	}

	return orderId, nil
}

func (r *repository) registerOrder(ctx context.Context) (uint64, error) {
	builder := sq.Insert(table.Orders).
		Columns("status").
		Values(schema.StatusAwaitingPayment).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING order_id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "repository.registerOrder",
		QueryRaw: query,
	}

	row, err := r.db.DB().Query(ctx, q, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	var orderId uint64
	if row.Next() {
		err = row.Scan(&orderId)
		if err != nil {
			return 0, err
		}
	}

	return orderId, nil
}

func (r *repository) saveOrderItems(ctx context.Context, orderId uint64, items []model.Item) error {
	builder := sq.Insert(table.OrdersItems).
		Columns("order_id", "item_id", "count").
		PlaceholderFormat(sq.Dollar)

	for _, item := range items {
		builder.Values(orderId, item.Sku, item.Count)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "repository.saveOrderItems",
		QueryRaw: query,
	}

	commandTag, err := r.db.DB().Exec(ctx, q, args...)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() < 1 {
		return libErr.ErrNoAddedRows
	}

	return nil
}

func (r *repository) saveOrderUser(ctx context.Context, orderId uint64, userId int64) error {
	builder := sq.Insert(table.OrdersUsers).
		Columns("order_id", "user_id").
		Values(orderId, userId).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "repository.saveOrderUser",
		QueryRaw: query,
	}

	commandTag, err := r.db.DB().Exec(ctx, q, args...)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() < 1 {
		return libErr.ErrNoAddedRows
	}

	return nil
}
