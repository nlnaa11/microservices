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

func (r *repository) RemoveFromStocks(ctx context.Context, itemId uint32, whIdToRemove []int64, whToUpdate *model.Stock) error {
	if len(whIdToRemove) == 0 {
		return errors.New("no warehouses indices to remove")
	}

	for _, whId := range whIdToRemove {
		err := r.removeFromStocks(ctx, itemId, whId)
		if err != nil {
			return errors.WithMessage(err, "removing warehouse entry from stocks")
		}
	}

	if whToUpdate != nil {
		err := r.updateInStocks(ctx, whToUpdate)
		if err != nil {
			return errors.WithMessage(err, "updating warehouse entry in stocks")
		}
	}

	return nil
}

func (r *repository) removeFromStocks(ctx context.Context, itemId uint32, whId int64) error {
	builder := sq.Delete(table.ItemsStocks).
		Where(sq.Eq{"warehouse_id": whId},
			sq.Eq{"item_id": itemId}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "repository.removeFromStocks",
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

func (r *repository) updateInStocks(ctx context.Context, wh *model.Stock) error {
	builder := sq.Update(table.ReservedItems).
		Set("count", wh.Count).
		Where(sq.Eq{"warehouse_id": wh.WarehouseId}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "repository.updateInStocks",
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
