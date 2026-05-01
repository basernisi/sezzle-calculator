package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	httpadapter "github.com/basernisi/sezzle-calculator/backend/internal/adapters/http"
	"github.com/basernisi/sezzle-calculator/backend/internal/ports"
)

type contextKey string

const claimsContextKey contextKey = "authClaims"

var ErrMissingBearerToken = errors.New("missing bearer token")

func RequireAuth(validator ports.TokenValidator, publicRoutes map[string]struct{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if _, ok := publicRoutes[r.Method+" "+r.URL.Path]; ok {
				next.ServeHTTP(w, r)
				return
			}

			token, err := extractBearerToken(r.Header.Get("Authorization"))
			if err != nil {
				httpadapter.WriteUnauthorized(w)
				return
			}

			claims, err := validator.ValidateToken(r.Context(), token)
			if err != nil {
				httpadapter.WriteUnauthorized(w)
				return
			}

			ctx := context.WithValue(r.Context(), claimsContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func extractBearerToken(headerValue string) (string, error) {
	parts := strings.SplitN(headerValue, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || strings.TrimSpace(parts[1]) == "" {
		return "", ErrMissingBearerToken
	}

	return strings.TrimSpace(parts[1]), nil
}
