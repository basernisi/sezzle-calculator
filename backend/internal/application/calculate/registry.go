package calculate

import (
	"strings"

	"github.com/basernisi/sezzle-calculator/backend/internal/domain/calculator"
)

type OperationRegistry struct {
	operations map[string]calculator.Operation
}

func NewOperationRegistry(operations ...calculator.Operation) OperationRegistry {
	registry := OperationRegistry{
		operations: make(map[string]calculator.Operation, len(operations)),
	}

	for _, operation := range operations {
		registry.operations[strings.ToLower(operation.Name())] = operation
	}

	return registry
}

func (r OperationRegistry) Get(name string) (calculator.Operation, error) {
	operation, ok := r.operations[strings.ToLower(strings.TrimSpace(name))]
	if !ok {
		return nil, calculator.ErrUnsupportedOperation
	}

	return operation, nil
}
