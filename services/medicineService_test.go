package service

import (
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

// func (m *MockDB) Create(ctx context.Context, user models.User) (models.User, error) {
// 	m.Called(user)
// 	return models.User{}, nil
// }

// func TestCreateMedicine(t *testing.T) {
// 	mkDB := new(MockDB)
// 	var rpo models.MedicineRepository
// 	srv := NewMedicineService(rpo)
// 	medicine := models.Medicine{
// 		Name: "Paracetamol",
// 		Dosage: "20mg",
// 		Description: "Nice med",
// 		Manufacturer: "Square",
// 		Price: 20,
// 		Stock: 100,
// 	}
// 	mkDB.On("Create", medicine).Return(medicine, nil)
// 	srv.CreateMedicine(c
// }
