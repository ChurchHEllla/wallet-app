package get_api_1_wallets

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

func (h *Handle) Handler(c *gin.Context) {

	idStr := c.Param("walletId")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing wallet id"})
		return
	}

	walletID, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("Received wallet ID: %s", idStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid wallet id"})
		return
	}

	wallet, err := h.service.GetWallet(c.Request.Context(), walletID)
	if err != nil {
		if err.Error() == "wallet not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := GetWalletResponse{
		ID:        wallet.ID.String(),
		Balance:   wallet.Balance,
		CreatedAt: wallet.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	c.JSON(http.StatusOK, resp)
}
