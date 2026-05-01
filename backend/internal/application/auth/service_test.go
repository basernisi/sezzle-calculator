package auth

import (
	"context"
	"errors"
	"testing"
)

type stubIssuer struct {
	token string
	err   error
}

func (s stubIssuer) IssueToken(ctx context.Context, subject string) (string, error) {
	if s.err != nil {
		return "", s.err
	}

	return s.token, nil
}

func TestIssueDemoToken(t *testing.T) {
	service := NewService("demo-client", "demo-secret", stubIssuer{token: "issued-token"})

	response, err := service.IssueDemoToken(context.Background(), TokenRequest{
		ClientID:     "demo-client",
		ClientSecret: "demo-secret",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if response.AccessToken != "issued-token" {
		t.Fatalf("expected issued token, got %q", response.AccessToken)
	}
}

func TestIssueDemoTokenInvalidCredentials(t *testing.T) {
	service := NewService("demo-client", "demo-secret", stubIssuer{token: "issued-token"})

	_, err := service.IssueDemoToken(context.Background(), TokenRequest{
		ClientID:     "wrong",
		ClientSecret: "demo-secret",
	})
	if !errors.Is(err, ErrInvalidClientCredentials) {
		t.Fatalf("expected invalid credentials error, got %v", err)
	}
}
