package handlers

// type Medicine struct {
// 	ID           int     `json:"id"`
// 	Name         string  `json:"name"`
// 	Dosage       string  `json:"dosage"`
// 	Manufacturer string  `json:"manufacturer"`
// 	Price        float64 `json:"price"`
// 	Stock        int     `json:"stock"`
// 	Created_at   time.Time
// 	Updated_at   time.Time
// }

// type MedicineApp struct {
// 	Router *gin.Engine
// }

// var medicines = []Medicine{}

// func (medApp MedicineApp) CreateMedicine(ctx *gin.Context) {
// 	var newMedicine Medicine
// 	if err := ctx.ShouldBindJSON(&newMedicine); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	medicines = append(medicines, newMedicine)
// 	ctx.JSON(http.StatusOK, newMedicine)
// }
