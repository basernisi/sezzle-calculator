package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jnsilvag/sezzle-calculator/backend/internal/ports"
)

type stubTokenValidator struct {
	claims ports.TokenClaims
	err    error
}

func (s stubTokenValidator) ValidateToken(ctx context.Context, token string) (ports.TokenClaims, error) {
	return s.claims, s.err
}

func TestAuthMiddleware(t *testing.T) {
	protected := RequireAuth(stubTokenValidator{claims: ports.TokenClaims{Subject: "user-123"}}, nil)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	request := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", nil)
	request.Header.Set("Authorization", "Bearer valid-token")
	responseRecorder := httptest.NewRecorder()

	protected.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, responseRecorder.Code)
	}
}

func TestAuthMiddlewareMissingToken(t *testing.T) {
	protected := RequireAuth(stubTokenValidator{}, nil)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	request := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", nil)
	responseRecorder := httptest.NewRecorder()

	protected.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, responseRecorder.Code)
	}
}

func TestAuthMiddlewareInvalidToken(t *testing.T) {
	protected := RequireAuth(stubTokenValidator{err: context.DeadlineExceeded}, nil)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	request := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", nil)
	request.Header.Set("Authorization", "Bearer invalid-token")
	responseRecorder := httptest.NewRecorder()

	protected.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, responseRecorder.Code)
	}
}

func TestAuthMiddlewareAllowsPublicRoute(t *testing.T) {
	protected := RequireAuth(stubTokenValidator{}, map[string]struct{}{
		"POST /api/v1/auth/token": {},
	})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/token", nil)
	responseRecorder := httptest.NewRecorder()

	protected.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, responseRecorder.Code)
	}
}
