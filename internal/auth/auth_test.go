package auth

import "testing"

func TestHashPassword(t *testing.T) {
	password := ""
	_, got := HashPassword(password)
	if got != errPassNotProvided {
		t.Errorf("wanted error, got %v", got)
	}
}

func getHashPassForTest(t *testing.T, pass string) string {
	t.Helper()
	hash, err := HashPassword(pass)
	if err != nil {
		t.Fatalf("failed to create test hash: %v", err)
	}

	return hash
}

func TestCheckHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		hash     string
		want error
	}{
		{
			name:     "empty password",
			password: "",
			hash:     "",
			want: errPassNotProvided,
		},
		{
			name:     "password validation",
			password: "5atWGC#$%",
			hash:     getHashPassForTest(t, "5atWGC#$%"),
			want: nil,
		},
		{
			name:     "Wrong password",
			password: "wrongItIs",
			hash:     getHashPassForTest(t, "wrongitis"),
			want: errWrongPass,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := CheckPasswordHash(tc.password, tc.hash)
			if err != tc.want {
				t.Errorf("Test %v - '%s' FAIL: \nexpected %v, \nactual: %v", i, tc.name, tc.want, err)
			}

		})
	}
}
