package ports

import "context"

type TokenIssuer interface {
	IssueToken(ctx context.Context, subject string) (string, error)
}
