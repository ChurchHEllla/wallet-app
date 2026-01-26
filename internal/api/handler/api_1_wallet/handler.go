package api_1_wallet

import (
	"log"
	"net/http"
	"wallet-app/internal/domain/wallet/models"
	"wallet-app/internal/domain/wallet/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateBalanceRequest struct {
	WalletID      string               `json:"walletId"`      // spelling
	OperationType models.OperationType `json:"operationType"` // "DEPOSIT" или "WITHDRAW"
	Amount        int64                `json:"amount"`
}

type UpdateBalanceResponse struct {
	Status string `json:"status"`
}

type Handle struct {
	service *service.Service
}

func New(s *service.Service) *Handle {
	return &Handle{service: s}
}

func (h *Handle) Handler(c *gin.Context) {
	var req UpdateBalanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	walletID, err := uuid.Parse(req.WalletID)
	if err != nil {
		log.Printf("Received wallet ID: %s", req.WalletID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid wallet id"})
		return
	}

	if !req.OperationType.Valid() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid operation type"})
		return
	}

	if req.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "amount must be positive"})
		return
	}

	err = h.service.UpdateBalance(c.Request.Context(), walletID, req.Amount, req.OperationType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, UpdateBalanceResponse{Status: "ok"})
}
