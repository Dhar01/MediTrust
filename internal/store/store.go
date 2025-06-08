package store

import (
	"context"
	"medicine-app/internal/database"
	"medicine-app/internal/repository"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines the unified access of repo layers
type Store struct {
	ProductRepo repository.ProductRepo
	MedRepo     repository.MedicineRepo
}

// NewStoreWithTx returns an instance of store with transaction support
func NewStoreWithTx(tx pgx.Tx) *Store {
	q := database.New(tx)

	return &Store{
		ProductRepo: repository.NewProdPostgresRepo(q),
		MedRepo:     repository.NewMedPostgresRepo(q),
	}
}

// NewStore returns an instance of store
func NewStore(q *database.Queries) *Store {
	return &Store{
		ProductRepo: repository.NewProdPostgresRepo(q),
		MedRepo:     repository.NewMedPostgresRepo(q),
	}
}

// WithTx is a helper function to handle database transactions
func WithTx(ctx context.Context, db *pgxpool.Pool, fn func(*Store) error) error {
	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	s := NewStoreWithTx(tx)

	if err := fn(s); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
