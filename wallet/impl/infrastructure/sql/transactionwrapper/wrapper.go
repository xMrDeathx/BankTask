package transactionwrapper

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TransactionWrapper struct {
	conn *pgxpool.Pool
}

func NewTransactionWrapper(conn *pgxpool.Pool) *TransactionWrapper {
	return &TransactionWrapper{conn}
}

func (t *TransactionWrapper) ExecuteWithTransaction(ctx context.Context, fn func(context.Context, pgx.Tx) error) error {
	tx, err := t.conn.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	err = fn(ctx, tx)
	if err != nil {
		rollbackErr := tx.Rollback(ctx)
		if rollbackErr != nil {
			return errors.New("rollback error: " + rollbackErr.Error())
		}
		return err
	}

	return nil
}
