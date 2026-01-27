package api_1_wallet

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"wallet-app/internal/api/handler/api_1_wallet/mocks"
	"wallet-app/internal/domain/wallet/models"

	"github.com/gin-gonic/gin"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_UpdateBalance_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mc := minimock.NewController(t)

	serviceMock := mocks.NewServiceMock(mc)
	handler := New(serviceMock)

	initialBalance := int64(500)
	walletID := uuid.New()
	opType := models.OperationType_Deposit
	reqBody := UpdateBalanceRequest{
		walletID.String(),
		opType,
		initialBalance,
	}

	bodyBytes, err := json.Marshal(reqBody)
	require.NoError(t, err)

	serviceMock.UpdateBalanceMock.
		Expect(minimock.AnyContext, walletID, initialBalance, opType).
		Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/wallets", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	handler.Handler(c)

	require.Equal(t, http.StatusOK, w.Code)

	var resp UpdateBalanceResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)

	assert.Equal(t, "ok", resp.Status)
}

func TestHandler_UpdateBalance_InvalidWalletID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mc := minimock.NewController(t)
	serviceMock := mocks.NewServiceMock(mc)
	handler := New(serviceMock)

	reqBody := UpdateBalanceRequest{
		WalletID:      "not-a-uuid",
		OperationType: models.OperationType_Deposit,
		Amount:        100,
	}

	bodyBytes, err := json.Marshal(reqBody)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/wallets", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	handler.Handler(c)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHandler_UpdateBalance_InvalidOperationType(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mc := minimock.NewController(t)
	serviceMock := mocks.NewServiceMock(mc)
	handler := New(serviceMock)

	reqBody := UpdateBalanceRequest{
		WalletID:      uuid.New().String(),
		OperationType: models.OperationType("123"), // невалидный
		Amount:        100,
	}

	bodyBytes, err := json.Marshal(reqBody)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/wallets", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	handler.Handler(c)

	require.Equal(t, http.StatusBadRequest, w.Code)
}
