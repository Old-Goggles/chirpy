package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestJWT(t *testing.T) {
	tokenSecret := "secret"
	userID := uuid.New()
	expiresIn := time.Hour

	token, err := MakeJWT(userID, tokenSecret, expiresIn)
	if err != nil {
		t.Fatalf("unexpected error making JWT: %v", err)
	}

	parsedID, err := ValidateJWT(token, tokenSecret)
	if err != nil {
		t.Fatalf("unexpected error validating JWT: %v", err)
	}

	if parsedID != userID {
		t.Errorf("expected ID %v, got %v", userID, parsedID)
	}
}

func TestExpiredJWT(t *testing.T) {
	tokenSecret := "secret"
	userID := uuid.New()
	expiresIn := -time.Hour

	token, err := MakeJWT(userID, tokenSecret, expiresIn)
	if err != nil {
		t.Fatalf("unexpected error making JWT: %v", err)
	}

	_, err = ValidateJWT(token, tokenSecret)
	if err == nil {
		t.Errorf("expected error for expired token, but got none")
	}
}

func TestWrongSecretJWT(t *testing.T) {
	userID := uuid.New()
	expiresIn := time.Hour

	token, _ := MakeJWT(userID, "correct-secret", expiresIn)

	_, err := ValidateJWT(token, "wrong-secret")
	if err == nil {
		t.Errorf("expected error for wrong secret, but got none")
	}
}
