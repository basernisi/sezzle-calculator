package middleware

import (
	"log/slog"
	"net/http"

	httpadapter "github.com/basernisi/sezzle-calculator/backend/internal/adapters/http"
)

func Recovery(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if recovered := recover(); recovered != nil {
					logger.Error("panic recovered", slog.Any("panic", recovered))
					httpadapter.WriteInternalError(w)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
