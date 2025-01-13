package models

import "context"

type GeneralService interface {
	ResetMedicineService(ctx context.Context) error
	ResetUserService(ctx context.Context) error
	ResetAddressService(ctx context.Context) error
}

type GeneralRepository interface {
	ResetMedicineRepo(ctx context.Context) error
	ResetUserRepo(ctx context.Context) error
	ResetAddressRepo(ctx context.Context) error
}
