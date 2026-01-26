package create_api_1_wallet

import (
	"net/http"
	"wallet-app/internal/domain/wallet/service"

	"github.com/gin-gonic/gin"
)

type CreateWalletRequest struct {
	InitialBalance int64 `json:"initial_balance"`
}

type CreateWalletResponse struct {
	ID        string `json:"id"`
	Balance   int64  `json:"balance"`
	CreatedAt string `json:"created_at"`
}

type Handle struct {
	service *service.Service
}

func New(s *service.Service) *Handle {
	return &Handle{service: s}
}

func (h *Handle) Handler(c *gin.Context) {
	var req CreateWalletRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	wallet, err := h.service.CreateWallet(c.Request.Context(), req.InitialBalance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := CreateWalletResponse{
		ID:        wallet.ID.String(),
		Balance:   wallet.Balance,
		CreatedAt: wallet.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	c.JSON(http.StatusCreated, resp)
}
