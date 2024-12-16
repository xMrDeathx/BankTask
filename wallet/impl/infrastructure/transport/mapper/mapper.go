package mapper

import (
	frontendapi "BankTask/wallet/api/frontend"
	"BankTask/wallet/impl/app/commands/walletcommand"
)

func MapBalanceToComponentBalance(balance walletcommand.BalanceResult) frontendapi.GetBalanceResponse {
	return frontendapi.GetBalanceResponse{
		Balance: balance.Balance,
	}
}
