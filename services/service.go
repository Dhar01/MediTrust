package service

import (
	"context"
	"medicine-app/models"
)

type generalService struct {
	Repo models.GeneralRepository
}

func NewGeneralService(genRepo models.GeneralRepository) models.GeneralService {
	return &generalService{
		Repo: genRepo,
	}
}

func (gs *generalService) ResetMedicineService(ctx context.Context) error {
	return gs.Repo.ResetMedicineRepo(ctx)
}

func (gs *generalService) ResetUserService(ctx context.Context) error {
	return gs.Repo.ResetUserRepo(ctx)
}

func (gs *generalService) ResetAddressService(ctx context.Context) error {
	return gs.Repo.ResetAddressRepo(ctx)
}
