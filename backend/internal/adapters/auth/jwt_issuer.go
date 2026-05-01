package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

type JWTIssuer struct {
	secret []byte
}

func NewJWTIssuer(secret string) JWTIssuer {
	return JWTIssuer{secret: []byte(secret)}
}

func (i JWTIssuer) IssueToken(_ context.Context, subject string) (string, error) {
	header, err := encodeJWTPart(map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	})
	if err != nil {
		return "", err
	}

	payload, err := encodeJWTPart(map[string]any{
		"sub": subject,
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	if err != nil {
		return "", err
	}

	unsignedToken := fmt.Sprintf("%s.%s", header, payload)
	mac := hmac.New(sha256.New, i.secret)
	mac.Write([]byte(unsignedToken))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return fmt.Sprintf("%s.%s", unsignedToken, signature), nil
}

func encodeJWTPart(value any) (string, error) {
	bytes, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(bytes), nil
}
