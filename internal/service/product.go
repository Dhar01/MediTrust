package service

import (
	"context"
	"medicine-app/internal/database/model"
	"medicine-app/internal/store"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MedService struct {
	db    *pgxpool.Pool
	store *store.Store
}

func NewMedService(db *pgxpool.Pool, store *store.Store) *MedService {
	return &MedService{
		db:    db,
		store: store,
	}
}

// CreateMedicine service creates a medicine and returns updated response
func (s *MedService) CreateMedicine(ctx context.Context, req model.MedicineRequest) (model.Response, error) {
	// ! check if the same name from same manufacturer with same dosage exists

	var result *model.Medicine

	err := store.WithTx(ctx, s.db, func(s *store.Store) error {
		product, err := s.ProductRepo.Create(ctx, *toProduct(req))
		if err != nil {
			return err
		}

		medicine, err := s.MedRepo.Create(ctx, *toMedicine(req))
		if err != nil {
			return err
		}

		medicine.BuildProduct(*product)
		result = medicine
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateMedicine service updates a medicine and returns updated response
func (s *MedService) UpdateMedicine(ctx context.Context, req model.MedicineRequest) (model.Response, error) {
	var result *model.Medicine

	err := store.WithTx(ctx, s.db, func(s *store.Store) error {
		product, err := s.ProductRepo.Update(ctx, *toProduct(req))
		if err != nil {
			return err
		}

		medicine, err := s.MedRepo.Update(ctx, *toMedicine(req))
		if err != nil {
			return err
		}

		medicine.BuildProduct(*product)
		result = medicine
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Delete deletes a product record in the database
func (s *MedService) Delete(ctx context.Context, id uuid.UUID) error {
	if err := s.store.ProductRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *MedService) FetchByID(ctx context.Context, productID uuid.UUID) (model.Response, error) {
	product, err := s.store.ProductRepo.FetchByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	medicine, err := s.store.MedRepo.FetchByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	medicine.BuildProduct(*product)

	return medicine, nil
}

func (s *MedService) FetchProducts(ctx context.Context) (model.Response, error) {
	return nil, nil
}

func toProduct(req model.MedicineRequest) *model.Product {
	return &model.Product{
		Name:         req.Name,
		Description:  req.Description,
		Manufacturer: req.Manufacturer,
		Price:        req.Price,
		Cost:         req.Cost,
		Stock:        req.Stock,
		Type:         req.Type,
	}
}

func toMedicine(req model.MedicineRequest) *model.Medicine {
	return &model.Medicine{
		Dosage:     req.Dosage,
		ExpiryDate: req.ExpiryDate,
	}
}
