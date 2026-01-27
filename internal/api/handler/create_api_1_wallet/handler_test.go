package create_api_1_wallet

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"wallet-app/internal/api/handler/create_api_1_wallet/mocks"
	"wallet-app/internal/domain/wallet/repository/mapper_objects"

	"github.com/gin-gonic/gin"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_CreateWallet_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mc := minimock.NewController(t)

	serviceMock := mocks.NewServiceMock(mc)
	handler := New(serviceMock)

	initialBalance := int64(500)

	reqBody := CreateWalletRequest{
		InitialBalance: initialBalance,
	}

	bodyBytes, err := json.Marshal(reqBody)
	require.NoError(t, err)

	walletID := uuid.New()
	createdAt := time.Now().UTC().Truncate(time.Second)

	returnedWallet := &mapper_objects.Wallet{
		ID:        walletID,
		Balance:   initialBalance,
		CreatedAt: createdAt,
	}

	serviceMock.CreateWalletMock.
		ExpectInitialBalanceParam2(initialBalance).
		Return(returnedWallet, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/wallets", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	handler.Handler(c)

	require.Equal(t, http.StatusCreated, w.Code)

	var resp CreateWalletResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)

	assert.Equal(t, walletID.String(), resp.ID)
	assert.Equal(t, initialBalance, resp.Balance)
	assert.Equal(t, createdAt.Format(time.RFC3339), resp.CreatedAt)
}

func TestHandler_CreateWallet_NegativeInitialBalance(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mc := minimock.NewController(t)

	serviceMock := mocks.NewServiceMock(mc)
	handler := New(serviceMock)

	reqBody := CreateWalletRequest{
		InitialBalance: -100,
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

	// опционально — проверка тела ответа
	var resp map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)

	assert.Equal(t, "amount must be positive", resp["error"])
}
