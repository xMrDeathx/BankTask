package transport

import (
	"BankTask/errs"
	frontendapi "BankTask/wallet/api/frontend"
	"BankTask/wallet/impl/app/commands/walletcommand"
	"BankTask/wallet/impl/domain/services"
	"BankTask/wallet/impl/infrastructure/transport/mapper"
	"encoding/json"
	"errors"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"io"
	"net/http"
)

type walletServer struct {
	walletService services.WalletService
}

func NewWalletServer(
	walletService services.WalletService,
) frontendapi.ServerInterface {
	return &walletServer{
		walletService: walletService,
	}
}

func (server *walletServer) ChangeBalance(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var balanceReq frontendapi.ChangeWalletRequest
	err = json.Unmarshal(reqBody, &balanceReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = validateReq(&balanceReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.walletService.ChangeBalance(
		r.Context(),
		walletcommand.ChangeBalanceCommand{
			WalletID:  balanceReq.WalletId,
			Operation: walletcommand.OperationType(balanceReq.OperationType),
			Amount:    balanceReq.Amount,
		})
	if errors.Is(err, errs.ErrWalletNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (server *walletServer) GetBalance(w http.ResponseWriter, r *http.Request, wallet_id openapi_types.UUID) {
	balance, err := server.walletService.GetBalance(r.Context(), wallet_id)
	if errors.Is(err, errs.ErrWalletNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	balanceJson := mapper.MapBalanceToComponentBalance(balance)
	response, err := json.Marshal(balanceJson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func validateReq(req *frontendapi.ChangeWalletRequest) error {
	if req.OperationType != walletcommand.DEPOSIT && req.OperationType != walletcommand.WITHDRAW {
		return errors.New("incorrect operation type")
	}
	if req.Amount < 0 {
		req.Amount *= -1
	}

	return nil
}
