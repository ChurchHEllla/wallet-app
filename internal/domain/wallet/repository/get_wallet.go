package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"wallet-app/internal/domain/wallet/repository/mapper_objects"

	"github.com/google/uuid"
)

// GetWallet возвращает кошелёк по UUID
func (r *Repository) GetWallet(ctx context.Context, id uuid.UUID) (*mapper_objects.Wallet, error) {
	var wallet mapper_objects.Wallet
	query := `SELECT id, balance, created_at FROM wallets WHERE id = $1`
	err := r.db.GetContext(ctx, &wallet, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // кошелек не найден
		}
		return nil, errors.New(fmt.Sprintf("get wallet: %w", err))
	}
	return &wallet, nil
}
