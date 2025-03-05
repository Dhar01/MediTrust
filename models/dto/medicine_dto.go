package dto

// CreateMedicineDTO represents the request body required for creating a new medicine
// @Description DTO for creating a new medicine
type CreateMedicineDTO struct {
	Name         string `json:"name" binding:"required" example:"Paracetamol"`
	Description  string `json:"description" binding:"required" example:"Pain reliever"`
	Dosage       string `json:"dosage" binding:"required" example:"500mg"`
	Manufacturer string `json:"manufacturer" binding:"required" example:"XZY Pharma"`
	Price        int32  `json:"price" binding:"required,min=0" example:"50" format:"int32"`
	Stock        int32  `json:"stock" binding:"required,min=0" example:"75" format:"int32"`
}

// UpdateMedicineDTO represents the request body required for updating a medicine
// @Description DTO for updating a medicine information
type UpdateMedicineDTO struct {
	Name         string `json:"name,omitempty" example:"Paracetamol"`
	Description  string `json:"description,omitempty" example:"Pain reliever"`
	Dosage       string `json:"dosage,omitempty" example:"500mg"`
	Manufacturer string `json:"manufacturer,omitempty" example:"XZY Pharma"`
	Price        *int32 `json:"price,omitempty" example:"50" format:"int32"`
	Stock        *int32 `json:"stock,omitempty" example:"75" format:"int32"`
}
