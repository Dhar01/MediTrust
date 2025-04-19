package services

import (
	"medicine-app/internal/repository"
)

type Services struct {
	MedService  MedService
	AuthService AuthService
	UserService UserService
}

type AuthService any
type UserService any

func NewServices(repo *repository.Repository) *Services {
	if repo == nil {
		panic("repository can't be nil")
	}

	return &Services{
		MedService: NewMedicineService(repo.MedRepo),
	}
}
