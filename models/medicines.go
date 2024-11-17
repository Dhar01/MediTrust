package models

import (
	"errors"
	"fmt"
)

var (
	medicines = make(map[int]Medicine)

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

/*
Endpoints plan

- POST /medicines
- GET /medicines
- PUT /medicines
- DELETE /medicines/{id}


Framework to be used: Gin
ORM: GORM
*/

func EntryMedicine(med Medicine) {
	medicines[med.ID] = med
	fmt.Println("Medicine entry creation was successful!")
}

func FindMedicine(id int) (Medicine, error) {
	med, ok := medicines[id]
	if !ok {
		return Medicine{}, medicineNotFound
	}
	return med, nil
}
