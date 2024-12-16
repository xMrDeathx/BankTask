package services

import (
	"BankTask/wallet/impl/app/commands/walletcommand"
	"BankTask/wallet/impl/app/mapper/walletmapper"
	"BankTask/wallet/impl/domain/model"
	"BankTask/wallet/impl/domain/repositories"
	"BankTask/wallet/impl/domain/services"
	"context"
	"github.com/google/uuid"
)

func NewWalletService(repository repositories.WalletRepository) services.WalletService {
	return &walletService{repository: repository}
}

type walletService struct {
	repository repositories.WalletRepository
}

func (service *walletService) GetBalance(ctx context.Context, walletID uuid.UUID) (walletcommand.BalanceResult, error) {
	balance, err := service.repository.GetBalance(ctx, walletID)
	if err != nil {
		return walletcommand.BalanceResult{}, err
	}

	return walletmapper.NewBalanceResultFromEntity(balance), nil
}

func (service *walletService) ChangeBalance(ctx context.Context, command walletcommand.ChangeBalanceCommand) error {
	currentBalance, err := service.repository.GetBalance(ctx, command.WalletID)
	if err != nil {
		return err
	}

	var newBalance int
	switch command.Operation {
	case walletcommand.DEPOSIT:
		newBalance = currentBalance + command.Amount
	case walletcommand.WITHDRAW:
		newBalance = currentBalance - command.Amount
	}

	wallet := model.Wallet{
		ID:      command.WalletID,
		Balance: newBalance,
	}
	return service.repository.Update(ctx, wallet)
}
