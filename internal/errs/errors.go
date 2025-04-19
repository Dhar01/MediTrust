package errs

import "errors"

var (
	ErrNotFound = errors.New("not found")

	ErrMedicineNotExist  = errors.New("medicine not exist")
	ErrMedicineExist     = errors.New("medicine with same ID/Name exists")
	ErrMedicineNotUpdate = errors.New("medicine can't be updated")

	ErrUserNotExist       = errors.New("user not exist")
	ErrUserNotAuthorized  = errors.New("user not authorized")
	ErrUserInactive       = errors.New("user not active")
	ErrEmailAlreadyExists = errors.New("email already being used")

	ErrInternalServer = errors.New("internal server error")
)
