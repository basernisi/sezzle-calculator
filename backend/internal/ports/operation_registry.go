package ports

import "github.com/jnsilvag/sezzle-calculator/backend/internal/domain/calculator"

type OperationRegistry interface {
	Get(name string) (calculator.Operation, error)
}
