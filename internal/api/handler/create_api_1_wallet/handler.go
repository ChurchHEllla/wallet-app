package create_api_1_wallet

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handle struct {
	service Service
}

func New(s Service) *Handle {
	return &Handle{service: s}
}

// Handler создает кошелек
func (h *Handle) Handler(c *gin.Context) {
	var req CreateWalletRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if req.InitialBalance <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "amount must be positive"})
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
