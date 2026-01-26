package repository

import (
	"context"
	"errors"
	"fmt"
	"wallet-app/internal/domain/wallet/repository/mapper_objects"

	_ "github.com/lib/pq"
)

// CreateWallet вставляет кошелёк
func (r *Repository) CreateWallet(ctx context.Context, op *mapper_objects.Wallet) error {
	query := `
        INSERT INTO wallets (id, balance, created_at)
        VALUES (:id, :balance, :created_at)
        ON CONFLICT (id) DO NOTHING
    `
	_, err := r.db.NamedExecContext(ctx, query, op)
	if err != nil {
		return errors.New(fmt.Sprintf("create wallet: %w", err))
	}
	return err
}
