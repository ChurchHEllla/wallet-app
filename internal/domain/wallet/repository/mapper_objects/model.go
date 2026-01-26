package mapper_objects

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	ID        uuid.UUID `db:"id"`
	Balance   int64     `db:"balance"`
	CreatedAt time.Time `db:"created_at"`
}

type WalletOperation struct {
	ID            uuid.UUID `db:"id"`
	WalletID      uuid.UUID `db:"wallet_id"`
	Amount        int64     `db:"amount"`
	OperationType string    `db:"operation_type"` // DEPOSIT | WITHDRAW
	CreatedAt     time.Time `db:"created_at"`
}
