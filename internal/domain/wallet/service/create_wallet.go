package service

import (
	"context"
	"time"
	"wallet-app/internal/domain/wallet/models"
	"wallet-app/internal/domain/wallet/repository/mapper"
	"wallet-app/internal/domain/wallet/repository/mapper_objects"

	"github.com/google/uuid"
)

// CreateWallet создаёт новый кошелек и возвращает его
func (s *Service) CreateWallet(ctx context.Context, initialBalance int64) (*mapper_objects.Wallet, error) {

	wallet := &models.Wallet{
		ID:        uuid.New(),
		Balance:   initialBalance,
		CreatedAt: time.Now(),
	}

	walletObj := mapper.ParseWallet(*wallet)
	if err := s.repository.CreateWallet(ctx, &walletObj); err != nil {
		return nil, err
	}

	return &walletObj, nil
}
