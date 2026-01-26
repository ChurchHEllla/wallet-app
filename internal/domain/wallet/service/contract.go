package service

import (
	"context"
	"wallet-app/internal/domain/wallet/repository"
	"wallet-app/internal/domain/wallet/repository/mapper_objects"

	"github.com/google/uuid"
)

type Repository interface {
	CreateOperation(ctx context.Context, op *mapper_objects.WalletOperation) error
	CreateWallet(ctx context.Context, op *mapper_objects.Wallet) error
	GetWallet(ctx context.Context, id uuid.UUID) (*mapper_objects.Wallet, error)
	UpdateBalance(ctx context.Context, id uuid.UUID, amount int64) error
	WithTx(ctx context.Context, fn func(ctx context.Context, repo repository.Repository) error) error
}
