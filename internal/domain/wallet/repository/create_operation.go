package repository

import (
	"context"
	"errors"
	"fmt"
	"wallet-app/internal/domain/wallet/repository/mapper_objects"
)

// CreateOperation сщздает операцию изменения баланса
func (r *Repository) CreateOperation(ctx context.Context, op *mapper_objects.WalletOperation) error {
	query := `
		INSERT INTO wallet_operations (
			id,
			wallet_id,
			amount,
			operation_type
		) VALUES (
			:id,
			:wallet_id,
			:amount,
			:operation_type
		)
	`

	_, err := r.db.NamedExecContext(ctx, query, op)
	if err != nil {
		return errors.New(fmt.Sprintf("create wallet operation: %w", err))
	}

	return err
}
