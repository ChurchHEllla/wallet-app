package get_api_1_wallets

type GetWalletResponse struct {
	ID        string `json:"id"`
	Balance   int64  `json:"balance"`
	CreatedAt string `json:"created_at"`
}
