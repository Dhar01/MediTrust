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

	// t.Run("testing MakeJWT", func(t *testing.T) {
	// 	_, got := GenerateAccessToken(userID, models.Admin, "", expiresIn)
	// 	want := errNoTokenProvided
	// 	if got != want {
	// 		t.Errorf("Expected %v, got %v", want, got)
	// 	}
	// })

	t.Run("testing AccessToken - OK", func(t *testing.T) {
		tokenString, err := GenerateAccessToken(userID, models.Customer, tokenST, expiresIn)
		assertMessage(t, err)

		id, _, err := ValidateAccessToken(tokenString, tokenST)
		assertMessage(t, err)

		if userID != id {
			t.Errorf("Expected UUIDs to be same; \ngot %v, original %v", id, userID)
		}
	})
	t.Run("testing VerificationToken - OK", func(t *testing.T) {
		tokenString, err := GenerateVerificationToken(userID, models.Customer, tokenST)
		assertMessage(t, err)

		id, err := ValidateVerificationToken(tokenString, tokenST)
		assertMessage(t, err)

		if userID != id {
			t.Errorf("Expected UUIDs to be same; \ngot %v, original %v", id, userID)
		}
	})
	t.Run("testing VerificationToken - Not OK", func(t *testing.T) {
		tokenString, err := GenerateVerificationToken(userID, models.Admin, "noSecret")
		assertMessage(t, err)

		_, err = ValidateVerificationToken(tokenString, tokenST)
		if err == nil {
			t.Fatalf("Expected error, but didn't get it")
		}
	})
}

func TestInputChecker(t *testing.T) {
	id := uuid.New()
	role := models.Customer
	tokenSecret := "myTokenSecret"
	expiresIn := time.Minute

	t.Run("inputChecker - Ok", func(t *testing.T) {
		err := inputChecker(id, role, tokenSecret, expiresIn)
		assertMessage(t, err)
	})
	t.Run("inputChecker - null ID", func(t *testing.T) {
		err := inputChecker(uuid.Nil, role, tokenSecret, expiresIn)
		assertErrorMessage(t, err, errNoUUIDProvided)
	})
	t.Run("inputChecker - null role", func(t *testing.T) {
		err := inputChecker(id, "", tokenSecret, expiresIn)
		assertErrorMessage(t, err, errNoRoleProvided)
	})
	t.Run("inputChecker - null Secret", func(t *testing.T) {
		err := inputChecker(id, role, "", expiresIn)
		assertErrorMessage(t, err, errNoSecretProvided)
	})
	t.Run("inputChecker - Expiration nil", func(t *testing.T) {
		err := inputChecker(id, role, tokenSecret, -1)
		assertErrorMessage(t, err, errNoTimeProvided)
	})
}

func assertMessage(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func assertErrorMessage(t testing.TB, err error, wanted error) {
	t.Helper()
	if err != wanted {
		t.Errorf("Expected %v, got %v", err, wanted)
	}
}
