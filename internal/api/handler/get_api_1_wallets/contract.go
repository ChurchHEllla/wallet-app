package get_api_1_wallets

import (
	"context"
	"wallet-app/internal/domain/wallet/repository/mapper_objects"

	"github.com/google/uuid"
)

type Service interface {
	GetWallet(ctx context.Context, id uuid.UUID) (*mapper_objects.Wallet, error)
}
