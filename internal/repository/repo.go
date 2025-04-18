package repository

import (
	"medicine-app/internal/database"
)

type Repository struct {
	MedRepo  MedicineRepository
	AuthRepo AuthRepository
	UserRepo UserRepository
}

type AuthRepository any
type UserRepository any

func NewRepository(db *database.DB) *Repository {
	if db == nil {
		panic("database can't be nil")
	}

	return &Repository{
		MedRepo: NewMedicineRepo(db.Medicine),
	}
}
