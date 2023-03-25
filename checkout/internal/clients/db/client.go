package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

var _ Client = (*client)(nil)

type Client interface {
	Close() error
	DB() *DB
}

type client struct {
	db        *DB
	closeFunc context.CancelFunc
}

func New(ctx context.Context, config *pgxpool.Config) (*client, error) {
	connect, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	_, cancel := context.WithCancel(ctx)

	return &client{
		db: &DB{
			pool: connect,
		},
		closeFunc: cancel,
	}, nil
}

func (c *client) Close() error {
	if c != nil {
		if c.closeFunc != nil {
			c.closeFunc()
		}
		if c.db != nil {
			c.db.pool.Close()
		}
	}

	return nil
}

func (c *client) DB() *DB {
	return c.db
}
