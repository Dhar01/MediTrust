package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"medicine-app/models"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	errNoSecretProvided   = errors.New("token secret not provided")
	errAuthHeaderNotFound = errors.New("authorization header not found")
	errNoRoleProvided     = errors.New("no role provided")
	errNoUUIDProvided     = errors.New("no user ID provided")
	errNoTimeProvided     = errors.New("time not provided")
)

type Claims struct {
	UserID uuid.UUID
	Role   string
	jwt.RegisteredClaims
}

// Generate Access Token
func GenerateAccessToken(userID uuid.UUID, role, tokenSecret string, expiresIn time.Duration) (string, error) {
	if err := inputChecker(userID, role, tokenSecret, expiresIn); err != nil {
		return wrapEmptyError(fmt.Errorf("GenerateAccessToken: %w", err))
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
		return wrapEmptyError(fmt.Errorf("MakeJWT: failed to sign token - %w", err))
	}

	return signedToken, nil
}

// Validate Access Token
func ValidateAccessToken(tokenString, tokenSecret string) (uuid.UUID, string, error) {
	token, err := getToken(tokenString, tokenSecret)
	if err != nil {
		return wrapUUIDError(fmt.Errorf("ValidateJWT: failed to parse token - %w", err))
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return wrapUUIDError(fmt.Errorf("ValidateJWT: %w", jwt.ErrTokenSignatureInvalid))
	}

	if claims.Issuer != models.CompanyName {
		return wrapUUIDError(fmt.Errorf("ValidateJWT: %w", jwt.ErrTokenInvalidIssuer))
	}

	return claims.UserID, claims.Role, nil
}

/*
Generate Verification Token - generates verification token to verify user.
By default, it expires in 10 minutes.
*/
func GenerateVerificationToken(userid uuid.UUID, role, tokenSecret string) (string, error) {
	if err := inputChecker(userid, role, tokenSecret, time.Minute); err != nil {
		return wrapEmptyError(fmt.Errorf("GenerateVerificationToken: %w", err))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID: userid,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    models.CompanyName,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
		},
	})

	signedToken, err := token.SignedString(([]byte(tokenSecret)))
	if err != nil {
		return wrapEmptyError(fmt.Errorf("GenerateVerificationToken: failed to sign token - %w", err))
	}

	return signedToken, nil
}

// Validate Verification Token
func ValidateVerificationToken(tokenString, tokenSecret string) (bool, error) {
	token, err := getToken(tokenString, tokenSecret)
	if err != nil {
		return false, fmt.Errorf("ValidateVerificationToken: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return false, fmt.Errorf("ValidateVerificationToken: %w", jwt.ErrSignatureInvalid)
	}

	if claims.Issuer != models.CompanyName {
		return false, fmt.Errorf("ValidateVerificationToken: %w", jwt.ErrTokenInvalidIssuer)
	}

	return true, nil
}

// Get Bearer Token - Authorization
func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return wrapEmptyError(fmt.Errorf("GetBearerToken: %w", errAuthHeaderNotFound))
	}

	tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	return tokenString, nil
}

// Generate Refresh Token
func GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return wrapEmptyError(fmt.Errorf("GenerateRefreshToken: failed to generate - %w", err))
	}

	return hex.EncodeToString(b), nil
}

func getToken(tokenString, tokenSecret string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(tokenSecret), nil
		},
	)

	if err != nil {
		return &jwt.Token{}, err
	}

	return token, nil
}

func inputChecker(id uuid.UUID, role, tokenSecret string, expiresIn time.Duration) error {
	if tokenSecret == "" {
		return errNoSecretProvided
	}

	if role == "" {
		return errNoRoleProvided
	}

	if id == uuid.Nil {
		return errNoUUIDProvided
	}

	if expiresIn < 0 {
		return errNoTimeProvided
	}

	return nil
}

func wrapUUIDError(err error) (uuid.UUID, string, error) {
	return uuid.Nil, "", err
}

func wrapEmptyError(err error) (string, error) {
	return "", err
}
