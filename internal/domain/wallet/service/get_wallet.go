package service

import (
	"context"
	"errors"
	"log"
	"wallet-app/internal/domain/wallet/repository/mapper_objects"

	"github.com/google/uuid"
)

// GetWallet возвращает кошелек по UUID
func (s *Service) GetWallet(ctx context.Context, id uuid.UUID) (*mapper_objects.Wallet, error) {
	log.Printf("Received wallet ID: %s", id)
	wallet, err := s.repository.GetWallet(ctx, id)
	if err != nil {
		return nil, err
	}

	if wallet == nil {
		return nil, errors.New("wallet not found")
	}
	return wallet, nil
}
