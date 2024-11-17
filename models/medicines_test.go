package models

import "testing"

func TestEntryMedicine(t *testing.T) {
	t.Run("medicine entry found", func(t *testing.T) {
		med := Medicine{
			ID:    1,
			Name:  "Paracetamol",
			Price: 50.0,
			Stock: 100,
		}
		

	})
}
