package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func New(
	db *sqlx.DB,
	tx *sqlx.Tx,
) *Repository {
	return &Repository{
		db: db,
		tx: tx,
	}
}
