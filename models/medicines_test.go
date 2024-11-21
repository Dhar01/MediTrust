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
	t.Run("update medicine information", func(t *testing.T) {
		store.EntryMedicine(med)

		updateMed := Medicine{
			ID:           1,
			Name:         "ParaUpdate",
			Dosage:       "1000mg",
			Manufacturer: "ABC Pharma",
			Price:        20,
			Stock:        150,
		}

		got := store.UpdateMedicine(1, updateMed)
		expectNoError(t, got)
	})
	t.Run("delete medicine", func(t *testing.T) {
		store.EntryMedicine(med)
		got := store.DeleteMedicine(med.ID)
		expectNoError(t, got)
	})
}

func expectNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("expected no error, got %v", got)
	}
}
