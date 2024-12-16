package repositories

import (
	"BankTask/wallet/impl/domain/model"
	"context"
	"github.com/google/uuid"
)

type WalletRepository interface {
	GetBalance(ctx context.Context, walletID uuid.UUID) (int, error)
	Update(ctx context.Context, wallet model.Wallet) error
}
