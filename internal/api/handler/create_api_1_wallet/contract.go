package create_api_1_wallet

import (
	"context"
	"wallet-app/internal/domain/wallet/repository/mapper_objects"
)

type Service interface {
	CreateWallet(ctx context.Context, initialBalance int64) (*mapper_objects.Wallet, error)
}
