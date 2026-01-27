package api_1_wallet

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handle struct {
	service Service
}

func New(s Service) *Handle {
	return &Handle{service: s}
}

// Handler Проводит операцию изменения баланса кошелька
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
		switch err.Error() {
		case "wallet not found":
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case "insufficient funds":
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, UpdateBalanceResponse{Status: "ok"})
}
