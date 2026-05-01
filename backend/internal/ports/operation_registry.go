package ports

import "github.com/basernisi/sezzle-calculator/backend/internal/domain/calculator"

type OperationRegistry interface {
	Get(name string) (calculator.Operation, error)
}
