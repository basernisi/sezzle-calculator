package httpadapter

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	applicationauth "github.com/jnsilvag/sezzle-calculator/backend/internal/application/auth"
)

func TestAuthHandlerIssueToken(t *testing.T) {
	handler := NewAuthHandler(applicationauth.NewService(
		"demo-client",
		"demo-secret",
		stubTokenIssuer{token: "issued-token"},
	))

	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/token", bytes.NewBufferString(`{"client_id":"demo-client","client_secret":"demo-secret"}`))
	responseRecorder := httptest.NewRecorder()

	handler.IssueToken(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, responseRecorder.Code)
	}
}

func TestAuthHandlerIssueTokenInvalidCredentials(t *testing.T) {
	handler := NewAuthHandler(applicationauth.NewService(
		"demo-client",
		"demo-secret",
		stubTokenIssuer{token: "issued-token"},
	))

	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/token", bytes.NewBufferString(`{"client_id":"demo-client","client_secret":"wrong-secret"}`))
	responseRecorder := httptest.NewRecorder()

	handler.IssueToken(responseRecorder, request)

	if responseRecorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, responseRecorder.Code)
	}
}

type stubTokenIssuer struct {
	token string
}

func (s stubTokenIssuer) IssueToken(ctx context.Context, subject string) (string, error) {
	return s.token, nil
}
