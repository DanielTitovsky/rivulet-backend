package app_postgres_transaction

import (
	"context"
	"fmt"

	app_postgres_pool "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres/pool"
	"github.com/jackc/pgx/v5"
)

type TransactionManager struct {
	pool app_postgres_pool.Pool
}

type txContextKey struct{}

func NewTransactionManager(pool app_postgres_pool.Pool) *TransactionManager {
	return &TransactionManager{
		pool: pool,
	}
}

func (t *TransactionManager) WithinTransaction(
	ctx context.Context,
	fn func(ctx context.Context) error,
) error {
	tx, err := t.pool.Begin(ctx)

	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer func() { tx.Rollback(ctx) }()

	ctx = contextWithTx(ctx, tx)

	if err := fn(ctx); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func contextWithTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, txContextKey{}, tx)
}

func FromContext(ctx context.Context) (pgx.Tx, error) {

	rawValues := ctx.Value(txContextKey{})

	tx, ok := rawValues.(pgx.Tx)

	if !ok {
		return nil, fmt.Errorf("transaction not found in context")
	}

	return tx, nil
}
