package errs

import "errors"

var (
	ErrWalletNotFound      = errors.New("wallet not found")
	ErrBalanceChangeFailed = errors.New("failed to change balance")
)
