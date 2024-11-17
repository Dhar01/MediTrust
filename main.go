package main

import (
	"fmt"
	"medicine-app/models"
)

func main() {
	med := models.Medicine{
		ID:    1,
		Name:  "Paracetamol",
		Price: 50.0,
		Stock: 100,
	}

	models.EntryMedicine(med)

	found, ok := models.FindMedicine(1)
	if !ok {
		fmt.Errorf("no medicine found")
	}

	fmt.Println("Medicine found:\n", found.Name)
}
