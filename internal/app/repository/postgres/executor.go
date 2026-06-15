package app_postgres_executor

import (
	"context"

	app_postgres_pool "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres/pool"
	app_postgres_transaction "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres/transaction"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type QueryExecutor interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

func GetQueryExecutor(ctx context.Context, pool app_postgres_pool.Pool) QueryExecutor {
	tx, err := app_postgres_transaction.FromContext(ctx)

	if err != nil {
		return pool
	}

	return tx
}
