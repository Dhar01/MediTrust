package models

import "testing"

func TestMedicines(t *testing.T) {
	t.Run("medicine entry found", func(t *testing.T) {
		store := NewMedicineStore()
		med := Medicine{
			ID:    1,
			Name:  "Paracetamol",
			Price: 50.0,
			Stock: 100,
		}
		got := store.EntryMedicine(med)
		want := store.FindMedicine(med.ID)

		if got != want {
			t.Errorf("medicine not found")
		}
	})
}
