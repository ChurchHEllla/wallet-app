package service

import (
	"context"
	"errors"
	"wallet-app/internal/domain/wallet/models"
	"wallet-app/internal/domain/wallet/repository"
	"wallet-app/internal/domain/wallet/repository/mapper"

	"github.com/google/uuid"
)

// UpdateBalance пополняет баланс кошелька
func (s *Service) UpdateBalance(ctx context.Context, walletID uuid.UUID, amount int64, opType models.OperationType) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	return s.repository.WithTx(ctx, func(ctx context.Context, txRepo repository.Repository) error {
		var delta int64

		switch opType {
		case "DEPOSIT":
			delta = amount
		case "WITHDRAW":
			delta = -amount
		}

		if err := txRepo.UpdateBalance(ctx, walletID, delta); err != nil {
			if err.Error() == "wallet not found" {
				return err
			}
			if err.Error() == "insufficient funds" {
				return err
			}
			return err
		}

		op := models.WalletOperation{
			ID:            uuid.New(),
			WalletID:      walletID,
			Amount:        amount,
			OperationType: opType,
		}

		opObj := mapper.ParseWalletOperation(op)
		if err := txRepo.CreateOperation(ctx, &opObj); err != nil {
			return err
		}

		return nil
	})
}
