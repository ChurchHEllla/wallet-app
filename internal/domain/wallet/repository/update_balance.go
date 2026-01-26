package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

// UpdateBalance изменяет баланс кошелька на указанную сумму по UUID кошелька
func (r *Repository) UpdateBalance(ctx context.Context, id uuid.UUID, delta int64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var currentBalance int64
	err = tx.QueryRowContext(ctx, "SELECT balance FROM wallets WHERE id=$1 FOR UPDATE", id).Scan(&currentBalance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("wallet not found")
		}
		return err
	}

	if currentBalance+delta < 0 {
		return errors.New("insufficient funds")
	}

	_, err = tx.ExecContext(ctx, "UPDATE wallets SET balance = balance + $1 WHERE id = $2", delta, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
