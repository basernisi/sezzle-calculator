package httpadapter

import "net/http"

func WriteUnauthorized(w http.ResponseWriter) {
	writeError(w, http.StatusUnauthorized, "UNAUTHORIZED", "A valid bearer token is required")
}

func WriteInternalError(w http.ResponseWriter) {
	writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "An unexpected error occurred")
}
