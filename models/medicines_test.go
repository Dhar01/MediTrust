package models

import (
	"errors"
	"testing"
)

func TestMedicines(t *testing.T) {
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
		expectNoError(t, got)

		// check if the medicine was added
		if _, ok := store.medicines[med.ID]; !ok {
			t.Errorf("expected medicine to be added, but it wasn't")
		}
	})
	t.Run("find existing medicine", func(t *testing.T) {
		store.EntryMedicine(med)
		got := store.FindMedicine(1)
		expectNoError(t, got)
	})
	t.Run("Find non-existent medicine", func(t *testing.T) {
		got := store.FindMedicine(4)
		if !errors.Is(got, errMedicineNotFound) {
			t.Errorf("expected error %v, got %v", errMedicineNotFound, got)
		}
	})
}

func expectNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("expected no error, got %v", got)
	}
}
