package httpadapter

import (
	"encoding/json"
	"errors"
	"net/http"

	applicationauth "github.com/jnsilvag/sezzle-calculator/backend/internal/application/auth"
)

type TokenService interface {
	IssueDemoToken(rctx any, request applicationauth.TokenRequest) (applicationauth.TokenResponse, error)
}

type AuthHandler struct {
	service applicationauth.Service
}

type IssueTokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type IssueTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

func NewAuthHandler(service applicationauth.Service) AuthHandler {
	return AuthHandler{service: service}
}

func (h AuthHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/auth/token", h.IssueToken)
}

func (h AuthHandler) IssueToken(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodyBytes)
	defer r.Body.Close()

	var requestBody IssueTokenRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&requestBody); err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_JSON", "Request body is invalid")
		return
	}

	response, err := h.service.IssueDemoToken(r.Context(), applicationauth.TokenRequest{
		ClientID:     requestBody.ClientID,
		ClientSecret: requestBody.ClientSecret,
	})
	if err != nil {
		switch {
		case errors.Is(err, applicationauth.ErrInvalidClientCredentials):
			writeError(w, http.StatusUnauthorized, "INVALID_CLIENT", "Client credentials are invalid")
		default:
			WriteInternalError(w)
		}
		return
	}

	writeJSON(w, http.StatusOK, IssueTokenResponse{
		AccessToken: response.AccessToken,
		TokenType:   response.TokenType,
		ExpiresIn:   response.ExpiresIn,
	})
}
