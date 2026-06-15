package app_postgres_pool

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Close()
	GetTimeout() time.Duration
}

type ConnectinPool struct {
	*pgxpool.Pool
	optTimeOut time.Duration
}

func NewConnectinPool(ctx context.Context, config Config) (*ConnectinPool, error) {
	connectinString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DataBase,
	)

	pgxconfig, err := pgxpool.ParseConfig(connectinString)
	if err != nil {
		return nil, fmt.Errorf("parse pgxconfig: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxconfig)

	if err != nil {
		return nil, fmt.Errorf("Create pgxpool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Ping pool: %w", err)
	}

	return &ConnectinPool{
		Pool:       pool,
		optTimeOut: config.Timeout,
	}, nil
}

func (p *ConnectinPool) GetTimeout() time.Duration {
	return p.optTimeOut
}
