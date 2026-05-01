package server

import (
	"log/slog"
	"net/http"

	httpadapter "github.com/jnsilvag/sezzle-calculator/backend/internal/adapters/http"
	"github.com/jnsilvag/sezzle-calculator/backend/internal/adapters/http/middleware"
	"github.com/jnsilvag/sezzle-calculator/backend/internal/ports"
)

func NewRouter(handler httpadapter.Handler, authHandler httpadapter.AuthHandler, validator ports.TokenValidator, logger *slog.Logger, allowedOrigin string) http.Handler {
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)
	authHandler.RegisterRoutes(mux)

	return chain(
		mux,
		middleware.Recovery(logger),
		middleware.Logging(logger),
		middleware.SecurityHeaders(),
		middleware.CORS(allowedOrigin),
		middleware.RequireAuth(validator, map[string]struct{}{
			"POST /api/v1/auth/token":    {},
			"OPTIONS /api/v1/auth/token": {},
		}),
	)
}

func chain(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	wrapped := handler
	for i := len(middlewares) - 1; i >= 0; i-- {
		wrapped = middlewares[i](wrapped)
	}

	return wrapped
}
