package main

import (
	"fmt"
	"medicine-app/models"
)

func main() {
	store := models.NewMedicineStore()

	med := models.Medicine{
		ID:    1,
		Name:  "Paracetamol",
		Price: 50.0,
		Stock: 100,
	}

	store.EntryMedicine(med)

	found, err := store.FindMedicine(1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Medicine found:\n", found.Name)
}
