package service

// type mockDB struct {
// 	deleteMedicineCalled bool
// 	deleteMedicineError  error
// 	deleteMedicineID     uuid.UUID
// }

// func (m *mockDB) DeleteMedicine(ctx context.Context, id uuid.UUID) error {
// 	m.deleteMedicineCalled = true
// 	m.deleteMedicineID = id
// 	return m.deleteMedicineError
// }

// func TestDeleteMedicine(t *testing.T) {
// 	t.Run("successful deletion", func(t *testing.T) {
// 		mDB := &mockDB{}
// 		service := &medicineService{
// 			DB: mDB,
// 		}
// 		testID := uuid.New()
// 		err := service.DeleteMedicine(context.Background(), testID)
// 		if err != nil {
// 			t.Errorf("expected no error, got %v", err)
// 		}
// 		if !mDB.deleteMedicineCalled {
// 			t.Error("DeleteMedicine was not called")
// 		}
// 		if mDB.deleteMedicineID != testID {
// 			t.Errorf("wrong ID was deleted, got %v, want %v", mDB.deleteMedicineID, testID)
// 		}
// 	})
// }
