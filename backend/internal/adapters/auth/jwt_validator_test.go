package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestJWTValidatorValidateToken(t *testing.T) {
	validator, err := NewJWTValidator("test-secret")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	token := signedToken(t, "test-secret", map[string]any{
		"sub": "user-123",
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	claims, err := validator.ValidateToken(context.Background(), token)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if claims.Subject != "user-123" {
		t.Fatalf("expected subject user-123, got %q", claims.Subject)
	}
}

func TestJWTValidatorValidateTokenExpired(t *testing.T) {
	validator, err := NewJWTValidator("test-secret")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	token := signedToken(t, "test-secret", map[string]any{
		"sub": "user-123",
		"exp": time.Now().Add(-time.Hour).Unix(),
	})

	_, err = validator.ValidateToken(context.Background(), token)
	if !errors.Is(err, ErrExpiredToken) {
		t.Fatalf("expected expired token error, got %v", err)
	}
}

func signedToken(t *testing.T, secret string, payload map[string]any) string {
	t.Helper()

	headerBytes, err := json.Marshal(map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	})
	if err != nil {
		t.Fatalf("marshal header: %v", err)
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("marshal payload: %v", err)
	}

	header := base64.RawURLEncoding.EncodeToString(headerBytes)
	body := base64.RawURLEncoding.EncodeToString(payloadBytes)
	unsigned := fmt.Sprintf("%s.%s", header, body)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(unsigned))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return fmt.Sprintf("%s.%s", unsigned, signature)
}
