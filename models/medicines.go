package models

import (
	"errors"
	"fmt"
)

var (
	medicineNotFound = errors.New("Medicine entry not found")
)

type Medicine struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Dosage       string  `json:"dosage"`
	Manufacturer string  `json:"manufacturer"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
}

type MedicineStore struct {
	medicines map[int]Medicine
}

func NewMedicineStore() *MedicineStore {
	return &MedicineStore{
		medicines: make(map[int]Medicine),
	}
}

func (ms *MedicineStore) EntryMedicine(med Medicine) {
	ms.medicines[med.ID] = med
	fmt.Println("Medicine entry creation was successful!")
}

func (ms *MedicineStore) FindMedicine(id int) (Medicine, error) {
	med, ok := ms.medicines[id]
	if !ok {
		return Medicine{}, medicineNotFound
	}
	return med, nil
}

func (ms *MedicineStore) UpdateMedicine(id int, med Medicine) error {
	return nil
}

func (ms *MedicineStore) DeleteMedicine(id int) error {
	return nil
}
