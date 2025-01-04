package medicine

// type MedicineService interface {
// 	CreateMedicine(ctx *gin.Context, medicine models.Medicine) error
// 	GetMedicineByID(id uuid.UUID) (models.Medicine, error)
// 	UpdateMedicineByID(id uuid.UUID) error
// 	DeleteMedicineByID(id uuid.UUID) error
// }

// type MedicineServiceImpl struct {
// 	DB *database.Queries
// }

// func NewMedicineService(db *database.Queries) *MedicineServiceImpl {
// 	return &MedicineServiceImpl{
// 		DB: db,
// 	}
// }

// func (msi *MedicineServiceImpl) CreateMedicine(ctx *gin.Context, medicine models.Medicine) error {
// 	_, err := msi.DB.CreateMedicine(ctx, database.CreateMedicineParams{
// 		Name:         medicine.Name,
// 		Dosage:       medicine.Dosage,
// 		Manufacturer: medicine.Manufacturer,
// 		Price:        medicine.Price,
// 		Stock:        medicine.Stock,
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
