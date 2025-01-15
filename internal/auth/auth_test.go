package auth

import "testing"

func TestHashPassword(t *testing.T) {
	password := ""
	_, got := HashPassword(password)
	if got != errPassNotProvided {
		t.Errorf("wanted error, got %v", got)
	}
}

func gethashPass(pass string) string {
	hash, _ := HashPassword(pass)
	return hash
}

func TestCheckHashPassword(t *testing.T) {
	tests := []struct {
		Name     string
		Password string
		Hash     string
		Expected error
	}{
		{
			Name:     "no input password",
			Password: "",
			Hash:     "",
			Expected: errPassNotProvided,
		},
		{
			Name: "Password validation",
			Password: "5atWGC#$%",
			Hash: gethashPass("5atWGC#$%"),
			Expected: nil,
		},
	}

	for i, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			err := CheckPasswordHash(tc.Password, tc.Hash)
			if err != tc.Expected {
				t.Errorf("Test %v - '%s' FAIL: \nexpected %v, \nactual: %v", i, tc.Name, tc.Expected, err)
			}

		})
	}
}
