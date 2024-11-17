package models

import (
	"errors"
	"fmt"
)

var (
	errMedicineNotFound  = errors.New("medicine entry not found")
	errDuplicateMedicine = errors.New("duplicate medicine ID")
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

func (ms *MedicineStore) EntryMedicine(med Medicine) error {
	_, err := ms.FindMedicine(med.ID)
	if err != nil {
		return errDuplicateMedicine
	}
	ms.medicines[med.ID] = med
	fmt.Println("Medicine entry created successfully!")
	return nil
}

func (ms *MedicineStore) FindMedicine(id int) (Medicine, error) {
	med, ok := ms.medicines[id]
	if !ok {
		return Medicine{}, errMedicineNotFound
	}
	return med, nil
}

func (ms *MedicineStore) UpdateMedicine(id int, updateMed Medicine) error {
	_, err := ms.FindMedicine(updateMed.ID)
	if err != nil {
		return errMedicineNotFound
	}

	ms.medicines[updateMed.ID] = updateMed
	fmt.Println("Medicine entry updated successfully!")
	return nil
}

func (ms *MedicineStore) DeleteMedicine(id int) error {
	return nil
}
