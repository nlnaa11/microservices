package transaction

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/clients/db"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/clients/db/key"
)

type Handler func(ctx context.Context) error

type Manager interface {
	RepeatableRead(ctx context.Context, fn Handler) error
	Serializable(ctx context.Context, fn Handler) error
}

type manager struct {
	transactor db.Transactor
}

func New(db db.Transactor) *manager {
	return &manager{
		transactor: db,
	}
}

func (m *manager) transaction(ctx context.Context, opts pgx.TxOptions, fn Handler) (err error) {
	tx, ok := ctx.Value(key.KeyTx).(pgx.Tx)
	if ok {
		return fn(ctx)
	}

	tx, err = m.transactor.BeginTx(ctx, opts)
	if err != nil {
		return errors.WithMessage(err, "beginning transaction")
	}

	ctx = context.WithValue(ctx, key.KeyTx, tx)

	defer func() {
		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.WithMessagef(err, "rollback: %s", errRollback.Error())
			}
		}

		if err == nil {
			err = tx.Commit(ctx)
			if err != nil {
				err = errors.WithMessage(err, "commiting transaction")
			}
		}
	}()

	if err = fn(ctx); err != nil {
		err = errors.WithMessage(err, "executing transaction")
	}

	return err
}

func (m *manager) RepeatableRead(ctx context.Context, fn Handler) error {
	opts := pgx.TxOptions{IsoLevel: pgx.RepeatableRead}
	return m.transaction(ctx, opts, fn)
}

func (m *manager) Serializable(ctx context.Context, fn Handler) error {
	opts := pgx.TxOptions{IsoLevel: pgx.Serializable}
	return m.transaction(ctx, opts, fn)
}
