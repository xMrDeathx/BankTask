package sql

import (
	"BankTask/errs"
	"BankTask/wallet/impl/domain/model"
	"BankTask/wallet/impl/domain/repositories"
	"BankTask/wallet/impl/infrastructure/sql/transactionwrapper"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewWalletRepository(conn *pgxpool.Pool) repositories.WalletRepository {
	return &walletRepository{conn: conn}
}

type walletRepository struct {
	conn *pgxpool.Pool
}

func (repo *walletRepository) GetBalance(ctx context.Context, walletID uuid.UUID) (int, error) {
	var balance int

	err := repo.conn.QueryRow(ctx, `
		SELECT balance
		FROM wallet
		where id = $1`, walletID).Scan(&balance)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, errs.ErrWalletNotFound
	}
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func (repo *walletRepository) Update(ctx context.Context, wallet model.Wallet) error {
	wrapper := transactionwrapper.NewTransactionWrapper(repo.conn)
	err := wrapper.ExecuteWithTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		_, err := tx.Exec(ctx, `
		UPDATE wallet
		SET balance = $1
		WHERE id = $2`,
			wallet.Balance, wallet.ID)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
