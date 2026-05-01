package auth

import (
	"context"
	"errors"
	"strings"

	"github.com/jnsilvag/sezzle-calculator/backend/internal/ports"
)

var ErrInvalidClientCredentials = errors.New("invalid client credentials")

type Service struct {
	expectedClientID     string
	expectedClientSecret string
	issuer               ports.TokenIssuer
}

func NewService(expectedClientID, expectedClientSecret string, issuer ports.TokenIssuer) Service {
	return Service{
		expectedClientID:     expectedClientID,
		expectedClientSecret: expectedClientSecret,
		issuer:               issuer,
	}
}

func (s Service) IssueDemoToken(ctx context.Context, request TokenRequest) (TokenResponse, error) {
	if strings.TrimSpace(request.ClientID) == "" || strings.TrimSpace(request.ClientSecret) == "" {
		return TokenResponse{}, ErrInvalidClientCredentials
	}

	if request.ClientID != s.expectedClientID || request.ClientSecret != s.expectedClientSecret {
		return TokenResponse{}, ErrInvalidClientCredentials
	}

	token, err := s.issuer.IssueToken(ctx, request.ClientID)
	if err != nil {
		return TokenResponse{}, err
	}

	return TokenResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   3600,
	}, nil
}
