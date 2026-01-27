package create_api_1_wallet

type CreateWalletRequest struct {
	InitialBalance int64 `json:"initial_balance"`
}

type CreateWalletResponse struct {
	ID        string `json:"id"`
	Balance   int64  `json:"balance"`
	CreatedAt string `json:"created_at"`
}
