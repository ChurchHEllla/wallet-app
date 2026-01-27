package api_1_wallet

import (
	"context"
	"wallet-app/internal/domain/wallet/models"

	"github.com/google/uuid"
)

type Service interface {
	UpdateBalance(ctx context.Context, walletID uuid.UUID, amount int64, opType models.OperationType) error
}
