package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/basernisi/sezzle-calculator/backend/internal/ports"
)

var (
	ErrInvalidToken         = errors.New("invalid token")
	ErrUnsupportedAlgorithm = errors.New("unsupported algorithm")
	ErrExpiredToken         = errors.New("token expired")
)

type JWTValidator struct {
	secret []byte
}

type jwtPayload struct {
	Subject   string `json:"sub"`
	ExpiresAt int64  `json:"exp"`
}

func NewJWTValidator(secret string) (JWTValidator, error) {
	if strings.TrimSpace(secret) == "" {
		return JWTValidator{}, fmt.Errorf("jwt secret is required")
	}

	return JWTValidator{secret: []byte(secret)}, nil
}

func (v JWTValidator) ValidateToken(_ context.Context, token string) (ports.TokenClaims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return ports.TokenClaims{}, ErrInvalidToken
	}

	headerPayload := parts[0] + "." + parts[1]
	signature, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return ports.TokenClaims{}, ErrInvalidToken
	}

	mac := hmac.New(sha256.New, v.secret)
	mac.Write([]byte(headerPayload))
	if !hmac.Equal(signature, mac.Sum(nil)) {
		return ports.TokenClaims{}, ErrInvalidToken
	}

	headerBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return ports.TokenClaims{}, ErrInvalidToken
	}

	var header struct {
		Algorithm string `json:"alg"`
	}
	if err := json.Unmarshal(headerBytes, &header); err != nil {
		return ports.TokenClaims{}, ErrInvalidToken
	}

	if header.Algorithm != "HS256" {
		return ports.TokenClaims{}, ErrUnsupportedAlgorithm
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return ports.TokenClaims{}, ErrInvalidToken
	}

	var payload jwtPayload
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return ports.TokenClaims{}, ErrInvalidToken
	}

	if payload.ExpiresAt > 0 && time.Now().Unix() >= payload.ExpiresAt {
		return ports.TokenClaims{}, ErrExpiredToken
	}

	return ports.TokenClaims{Subject: payload.Subject}, nil
}
