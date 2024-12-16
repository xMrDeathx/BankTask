package walletmapper

import (
	"BankTask/wallet/impl/app/commands/walletcommand"
)

func NewBalanceResultFromEntity(balance int) walletcommand.BalanceResult {
	return walletcommand.BalanceResult{
		Balance: balance,
	}
}
