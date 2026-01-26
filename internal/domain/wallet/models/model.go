package models

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	ID        uuid.UUID
	Balance   int64
	CreatedAt time.Time
}

type WalletOperation struct {
	ID            uuid.UUID
	WalletID      uuid.UUID
	Amount        int64
	OperationType OperationType // DEPOSIT | WITHDRAW
	CreatedAt     time.Time
}
