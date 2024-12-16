package services

import (
	"BankTask/wallet/impl/app/commands/walletcommand"
	"context"
	"github.com/google/uuid"
)

type WalletService interface {
	GetBalance(ctx context.Context, walletID uuid.UUID) (walletcommand.BalanceResult, error)
	ChangeBalance(ctx context.Context, command walletcommand.ChangeBalanceCommand) error
}
