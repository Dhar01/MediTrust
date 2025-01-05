package controllers

import (
	"encoding/json"
	"medicine-app/internal/database"
	"medicine-app/models"
	"net/http"
)

type Config struct {
	DB       *database.Queries
	Platform string
}

func (cfg *Config) CreateMedicineHandler(w http.ResponseWriter, r *http.Request) {
	methodChecker(w, r, http.MethodPost)

	var newMedicine models.MedicineBody

	if err := json.NewDecoder(r.Body).Decode(&newMedicine); err != nil {
		respondWithError(w, http.StatusBadRequest, "can't decode request", err)
		return
	}

	defer r.Body.Close()

	medicine, err := cfg.DB.CreateMedicine(r.Context(), database.CreateMedicineParams{
		Name:         newMedicine.Name,
		Dosage:       newMedicine.Dosage,
		Manufacturer: newMedicine.Manufacturer,
		Price:        newMedicine.Price,
		Stock:        newMedicine.Stock,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "can't create medicine", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, models.Medicine{
		Name:         medicine.Name,
		Dosage:       medicine.Dosage,
		Manufacturer: medicine.Manufacturer,
		Price:        medicine.Price,
		Stock:        medicine.Stock,
	})
}

func (cfg *Config) DeleteMedicine(w http.ResponseWriter, r *http.Request) {
	methodChecker(w, r, http.MethodDelete)

	var medID models.MedicineID

	if err := json.NewDecoder(r.Body).Decode(&medID); err != nil {
		respondWithError(w, http.StatusBadRequest, "can't decode request", err)
		return
	}

	defer r.Body.Close()

	if err := cfg.DB.DeleteMedicine(r.Context(), medID.ID); err != nil {
		respondWithError(w, http.StatusInternalServerError, "can't delete medicine", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write([]byte("Deleted")); err != nil {
		respondWithError(w, http.StatusInternalServerError, "write error", err)
		return
	}
}