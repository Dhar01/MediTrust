package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"medicine-app/models"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	errNoTokenProvided    = errors.New("token secret not provided")
	errAuthHeaderNotFound = errors.New("authorization header not found")
	errNoRoleProvided     = errors.New("no role found")
	errNoUUIDProvided     = errors.New("no user ID provided")
)

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
	jwt.RegisteredClaims
}

// Generate Access Token
func MakeJWT(userID uuid.UUID, role, tokenSecret string, expiresIn time.Duration) (string, error) {
	if tokenSecret == "" {
		return wrapEmptyError(errNoTokenProvided)
	}

	if role == "" {
		return wrapEmptyError(errNoRoleProvided)
	}

	if userID == uuid.Nil {
		return wrapEmptyError(errNoUUIDProvided)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    models.CompanyName,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		},
	})

	signedToken, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return wrapEmptyError(err)
	}

	return signedToken, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, string, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(tokenSecret), nil
		},
	)
	if err != nil {
		return wrapUUIDError(err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return wrapUUIDError(jwt.ErrTokenSignatureInvalid)
	}

	if claims.Issuer != models.CompanyName {
		return wrapUUIDError(jwt.ErrTokenInvalidIssuer)
	}

	return claims.UserID, claims.Role, nil
}

func GetBearerToken(headers http.Header) (string, error) {
	tokenString, err := getHeader(headers, "Bearer")
	if err != nil {
		return wrapEmptyError(err)
	}

	return tokenString, nil
}

// Generate Refresh Token
func MakeRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return wrapEmptyError(err)
	}

	return hex.EncodeToString(b), nil
}

func getHeader(headers http.Header, key string) (string, error) {
	authHeader := headers.Get("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, key) {
		return wrapEmptyError(errAuthHeaderNotFound)
	}

	tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, key))
	return tokenString, nil
}

func wrapUUIDError(err error) (uuid.UUID, string, error) {
	return uuid.Nil, "", err
}

func wrapEmptyError(err error) (string, error) {
	return "", err
}
