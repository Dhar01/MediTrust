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
	if _, ok := ms.medicines[med.ID]; ok {
		return errDuplicateMedicine
	}

	// if _, ok := ms.medicines[med.Name]; ok {
	// 	return errDuplicateMedicine
	// }

	ms.medicines[med.ID] = med
	fmt.Println("Medicine entry created successfully!")
	return nil
}

func (ms *MedicineStore) FindMedicine(id int) error {
	if _, ok := ms.medicines[id]; !ok {
		return errMedicineNotFound
	}

	return nil
}

func (ms *MedicineStore) UpdateMedicine(id int, updateMed Medicine) error {
	if _, ok := ms.medicines[id]; !ok {
		return errMedicineNotFound
	}

	ms.medicines[id] = updateMed
	fmt.Println("Medicine entry updated successfully!")
	return nil
}

func (ms *MedicineStore) DeleteMedicine(medID int) error {
	if _, ok := ms.medicines[medID]; !ok {
		return errMedicineNotFound
	}

	delete(ms.medicines, medID)
	fmt.Println("Medicine entry deleted successfully!")
	return nil
}
