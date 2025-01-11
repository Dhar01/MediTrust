package service

// type MockDB struct {
// 	mock.Mock
// }

// func (m *MockDB) CreateMedicine(ctx context.Context, params database.CreateMedicineParams) (database.Medicine, error) {
// 	args := m.Called(ctx, params)
// 	return args.Get(0).(database.Medicine), args.Error(1)
// }

// func TestCreateMedicine(t *testing.T) {
// 	mockDB := new(MockDB)
// 	medService := NewMedicineService(*mockDB)
// 	mockDB.On("CreateMedicine", mock.Anything, database.CreateMedicineParams{
// 		Name:         "Aspirin",
// 		Dosage:       "500mg",
// 		Description:  "Pain Relief",
// 		Manufacturer: "PharmaCorp",
// 		Price:        100,
// 		Stock:        50,
// 	}).Return(database.Medicine{
// 		ID:           uuid.New(),
// 		Name:         "Aspirin",
// 		Dosage:       "500mg",
// 		Description:  "Pain Relief",
// 		Manufacturer: "PharmaCorp",
// 		Price:        100,
// 		Stock:        50,
// 		CreatedAt:    time.Now(),
// 		UpdatedAt:    time.Now(),
// 	}, nil)
// }
