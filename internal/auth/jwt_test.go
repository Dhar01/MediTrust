package auth

import (
	"medicine-app/models"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestJWT(t *testing.T) {

	// tests := []struct {
	// 	userID uuid.UUID,
	// 	tokenSecret string,
	// 	expiresIn time.Duration
	// } {
	// 	{
	// 		userID: uuid.New(),
	// 		tokenSecret: "",
	// 		expiresIn: time.Minute,
	// 	},
	// }

	userID := uuid.New()
	tokenST := "mysecretkey"
	expiresIn := time.Minute

	t.Run("testing MakeJWT", func(t *testing.T) {
		_, got := MakeJWT(userID, models.Admin, "", expiresIn)

		want := errNoTokenProvided
		if got != want {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})
	t.Run("testing ValidateJWT - OK", func(t *testing.T) {
		tokenString, err := MakeJWT(userID, models.Customer, tokenST, expiresIn)
		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}

		id, err := ValidateJWT(tokenString, tokenST)
		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}

		if userID != id {
			t.Errorf("Expected UUIDs to be same; \ngot %v, original %v", id, userID)
		}
	})
}
