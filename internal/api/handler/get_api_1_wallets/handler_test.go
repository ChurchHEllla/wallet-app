package get_api_1_wallets

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"wallet-app/internal/api/handler/get_api_1_wallets/mocks"
	"wallet-app/internal/domain/wallet/repository/mapper_objects"

	"github.com/gin-gonic/gin"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetWallet_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mc := minimock.NewController(t)

	serviceMock := mocks.NewServiceMock(mc)
	handler := New(serviceMock)

	walletID := uuid.New()
	expectedWallet := &mapper_objects.Wallet{
		ID:      walletID,
		Balance: 1000,
	}

	serviceMock.GetWalletMock.
		ExpectIdParam2(walletID).
		Return(expectedWallet, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets/"+walletID.String(), nil)
	c.Request = req

	c.Params = gin.Params{
		{Key: "walletId", Value: walletID.String()},
	}

	handler.Handler(c)

	require.Equal(t, http.StatusOK, w.Code)

	var resp mapper_objects.Wallet
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)

	assert.Equal(t, *expectedWallet, resp)
}

func TestHandler_GetWallet_MissingWalletID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mc := minimock.NewController(t)
	serviceMock := mocks.NewServiceMock(mc)
	handler := New(serviceMock)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets/", nil)
	c.Request = req

	// params не задаём вообще
	handler.Handler(c)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHandler_GetWallet_InvalidWalletID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mc := minimock.NewController(t)
	serviceMock := mocks.NewServiceMock(mc)
	handler := New(serviceMock)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets/not-a-uuid", nil)
	c.Request = req

	c.Params = gin.Params{
		{Key: "walletId", Value: "not-a-uuid"},
	}

	handler.Handler(c)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHandler_GetWallet_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mc := minimock.NewController(t)
	serviceMock := mocks.NewServiceMock(mc)
	handler := New(serviceMock)

	walletID := uuid.New()

	serviceMock.GetWalletMock.
		ExpectIdParam2(walletID).
		Return(nil, errors.New("wallet not found"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets/"+walletID.String(), nil)
	c.Request = req

	c.Params = gin.Params{
		{Key: "walletId", Value: walletID.String()},
	}

	handler.Handler(c)

	require.Equal(t, http.StatusNotFound, w.Code)
}
