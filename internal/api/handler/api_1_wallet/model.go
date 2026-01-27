package api_1_wallet

import "wallet-app/internal/domain/wallet/models"

type UpdateBalanceRequest struct {
	WalletID      string               `json:"walletId"`      // spelling
	OperationType models.OperationType `json:"operationType"` // "DEPOSIT" или "WITHDRAW"
	Amount        int64                `json:"amount"`
}

type UpdateBalanceResponse struct {
	Status string `json:"status"`
}
