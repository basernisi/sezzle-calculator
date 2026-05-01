package ports

import "context"

type TokenClaims struct {
	Subject string
}

type TokenValidator interface {
	ValidateToken(ctx context.Context, token string) (TokenClaims, error)
}
