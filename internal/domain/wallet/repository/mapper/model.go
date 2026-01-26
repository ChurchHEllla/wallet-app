package mapper

import (
	"wallet-app/internal/domain/wallet/models"
	"wallet-app/internal/domain/wallet/repository/mapper_objects"
)

// Безопасная и контролируемая конвертация типов данных

func ParseWallet(p models.Wallet) mapper_objects.Wallet {
	return mapper_objects.Wallet{
		ID:        p.ID,
		CreatedAt: p.CreatedAt,
		Balance:   p.Balance,
	}
}

func ParseWalletOperation(p models.WalletOperation) mapper_objects.WalletOperation {
	return mapper_objects.WalletOperation{
		ID:            p.ID,
		WalletID:      p.WalletID,
		Amount:        p.Amount,
		OperationType: string(p.OperationType),
		CreatedAt:     p.CreatedAt,
	}
}
