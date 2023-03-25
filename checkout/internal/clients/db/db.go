package db

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/db/key"
)

type Query struct {
	Name     string
	QueryRaw string
}

type SQLExecer interface {
	Get(ctx context.Context, dst interface{}, query Query, args ...interface{}) error
	Select(ctx context.Context, dst interface{}, query Query, args ...interface{}) error

	Exec(ctx context.Context, query Query, args ...interface{}) (pgconn.CommandTag, error)
	// Query(ctx context.Context, quety Query, args ...interface{}) (pgx.Rows, error)
	// QueryRow(ctx context.Context, query Query, args ...interface{}) pgx.Row
}

type Transactor interface {
	BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error)
}

var _ SQLExecer = (*DB)(nil)
var _ Transactor = (*DB)(nil)

type DB struct {
	pool *pgxpool.Pool
}

/// SQLEXECER

// обертка на pgxscan.Get
func (db *DB) Get(ctx context.Context, dst interface{}, query Query, args ...interface{}) error {
	return pgxscan.Get(ctx, db.pool, dst, query.QueryRaw, args...)
}

// обертка на pgxscan.Select
func (db *DB) Select(ctx context.Context, dst interface{}, query Query, args ...interface{}) error {
	return pgxscan.Select(ctx, db.pool, dst, query.QueryRaw, args...)
}

// обертка на pgxpool.Exec
func (db *DB) Exec(ctx context.Context, query Query, args ...interface{}) (pgconn.CommandTag, error) {
	tx, ok := ctx.Value(key.KeyTx).(pgx.Tx)
	if ok {
		return tx.Exec(ctx, query.QueryRaw, args...)
	}

	return db.pool.Exec(ctx, query.QueryRaw, args...)
}

/// TRANSACTIONS

// обертка на pgx.BeginTx
func (db *DB) BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error) {
	return db.pool.BeginTx(ctx, opts)
}
