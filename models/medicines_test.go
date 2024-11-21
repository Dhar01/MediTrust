package models

import (
	"errors"
	"testing"
)

func TestEntryMedicine(t *testing.T) {
	store := NewMedicineStore()

	med := Medicine{
		ID:           1,
		Name:         "Paracetamol",
		Dosage:       "500mg",
		Manufacturer: "XYZ Pharma",
		Price:        10.5,
		Stock:        100,
	}

	t.Run("add new medicine", func(t *testing.T) {
		got := store.EntryMedicine(med)
		if got != nil {
			t.Errorf("expected no error, got %v", got)
		}

		// check if the medicine was added
		if _, ok := store.medicines[med.ID]; !ok {
			t.Errorf("expected medicine to be added, but it wasn't")
		}
	})

}

func TestFindMedicines(t *testing.T) {
	store := NewMedicineStore()

	med := Medicine{
		ID:           1,
		Name:         "Paracetamol",
		Dosage:       "500mg",
		Manufacturer: "XYZ Pharma",
		Price:        10.5,
		Stock:        100,
	}

	store.EntryMedicine(med)

	t.Run("find existing medicine", func(t *testing.T) {
		got := store.FindMedicine(1)
		if got != nil {
			t.Errorf("expected no error, got %v", got)
		}
	})
	t.Run("Find non-existent medicine", func(t *testing.T) {
		got := store.FindMedicine(4)
		if !errors.Is(got, errMedicineNotFound) {
			t.Errorf("expected error %v, got %v", errMedicineNotFound, got)
		}
	})
}
