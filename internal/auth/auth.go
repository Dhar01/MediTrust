package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	errPassNotProvided = errors.New("password not provided")
	errWrongPass = errors.New("password hash not matching")
)

func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errPassNotProvided
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CheckPasswordHash(password, hash string) error {
	if password == "" || hash == "" {
		return errPassNotProvided
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return errWrongPass
	}

	return nil
}
