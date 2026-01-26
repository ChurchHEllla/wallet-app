package repository

import "context"

// WithTx обеспечивает безопасную транзакцию
func (r *Repository) WithTx(
	ctx context.Context,
	fn func(ctx context.Context, repo Repository) error,
) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			err = tx.Rollback()
			if err != nil {
				panic(err)
			}
			panic(p)
		}
	}()

	txRepo := New(r.db, tx)

	if err := fn(ctx, *txRepo); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
