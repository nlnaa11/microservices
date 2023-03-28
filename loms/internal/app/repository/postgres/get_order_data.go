package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	libErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/model"
	convert "gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/repository/postgres/converter"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/repository/postgres/schema"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/repository/postgres/table"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/clients/db"
)

func (r *repository) GetOrderData(ctx context.Context, orderId uint64) (*model.OrderInfo, error) {
	// from orders: status
	order, err := r.getOrderInfo(ctx, orderId)
	if err != nil {
		return nil, errors.WithMessage(err, "getting order info")
	}

	// from users_orders: userId
	user, err := r.getOrderUser(ctx, orderId)
	if err != nil {
		return nil, errors.WithMessage(err, "getting order user")
	}

	// from orders_items: items
	items, err := r.getOrderItems(ctx, orderId)
	if err != nil {
		return nil, errors.WithMessage(err, "getting order items")
	}

	orderData := convert.ToModelOrderInfo(items, order, user)

	return orderData, nil
}

func (r *repository) GetOrderInfo(ctx context.Context, orderId uint64) (*model.Order, error) {
	order, err := r.getOrderInfo(ctx, orderId)
	if err != nil {
		return nil, errors.WithMessage(err, "getting order info")
	}

	return &model.Order{
		OrderId: order.Id,
		Status:  order.Status.String(),
	}, nil
}

func (r *repository) getOrderInfo(ctx context.Context, orderId uint64) (*schema.Order, error) {
	builder := sq.Select(
		"order_id", "status").
		From(table.Orders).
		Where(sq.Eq{"order_id": orderId}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "repository.getOrderInfo",
		QueryRaw: query,
	}

	// TODO: проверить преобразование типа int8 в тип Status
	var order schema.Order
	if err := r.db.DB().Select(ctx, &order, q, args...); err != nil {
		return &order, err
	}

	return &order, nil
}

func (r *repository) getOrderUser(ctx context.Context, orderId uint64) (int64, error) {
	builder := sq.Select("user_id").
		From(table.OrdersUsers).
		Where(sq.Eq{"order_id": orderId}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "repository.getOrderUser",
		QueryRaw: query,
	}

	var user int64
	if err := r.db.DB().Select(ctx, &user, q, args...); err != nil {
		return 0, err
	}

	return user, nil
}

func (r *repository) GetOrderItems(ctx context.Context, orderId uint64) ([]model.Item, error) {
	items, err := r.getOrderItems(ctx, orderId)
	if err != nil {
		return nil, errors.WithMessage(err, "getting orders items")
	}

	result := convert.ToModelItems(items)
	if len(result) == 0 {
		return nil, libErr.ErrEmptyItems
	}

	return result, nil
}

func (r *repository) getOrderItems(ctx context.Context, orderId uint64) ([]schema.Item, error) {
	builder := sq.Select(
		"item_id, " + "count").
		From(table.Orders).
		Where(sq.Eq{"order_id": orderId}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "repository.getOrderItems",
		QueryRaw: query,
	}

	var items []schema.Item
	if err := r.db.DB().Select(ctx, &items, q, args...); err != nil {
		return nil, err
	}

	return items, nil
}
