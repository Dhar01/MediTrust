package repository

import (
	"errors"
	"medicine-app/internal/database"
	"medicine-app/internal/database/model"
	"medicine-app/internal/errs"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgconn"
)

// toDBProductType helper mapped the type to the database.ProductType
func toDBProductType(t model.ProductType) database.ProductType {
	return database.ProductType(t)
}

// toModelProductType helper mapped the type to the model.ProductType
func toModelProductType(t database.ProductType) model.ProductType {
	return model.ProductType(t)
}

// setErrorMsg sets the error based on error type
func setErrorMsg(err error) error {
	var pgErr *pgconn.PgError

	if errors.Is(err, pgx.ErrNoRows) {
		return errs.ErrNotFound
	}

	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return errs.ErrConflict
	}

	return err
}
