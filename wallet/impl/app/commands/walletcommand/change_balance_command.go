package walletcommand

import "github.com/google/uuid"

type ChangeBalanceCommand struct {
	WalletID  uuid.UUID
	Operation OperationType
	Amount    int
}

type OperationType string

const (
	DEPOSIT  = "DEPOSIT"
	WITHDRAW = "WITHDRAW"
)
