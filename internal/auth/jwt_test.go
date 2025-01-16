package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeJWT(t *testing.T) {
	userID := uuid.New()
	tokenST := ""
	expiresIn := time.Minute

	_, got := MakeJWT(userID, tokenST, expiresIn)
	want := errNoTokenProvided

	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
