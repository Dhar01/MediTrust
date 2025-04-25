package errs

import "errors"

var (
	ErrNotFound = errors.New("not found")

	ErrMedicineNotExist = errors.New("medicine not exist")
	ErrMedicineExist    = errors.New("medicine with same ID/Name exists")
	ErrMedicineUpdate   = errors.New("medicine can't be updated")

	ErrUserNotExist       = errors.New("user not exist")
	ErrUserNotAuthorized  = errors.New("user not authorized")
	ErrUserInactive       = errors.New("user not active")
	ErrUserRoleNotFound   = errors.New("user role not found")
	ErrUserIdNotFound     = errors.New("user ID not found")
	ErrEmailAlreadyExists = errors.New("email already being used")

	ErrInternalServer = errors.New("internal server error")

	ErrUserNotAdmin = errors.New("role is not admin")
)
