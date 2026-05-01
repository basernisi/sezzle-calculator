package calculate

import (
	"errors"
	"testing"

	"github.com/jnsilvag/sezzle-calculator/backend/internal/domain/calculator"
)

type stubRegistry struct {
	operation calculator.Operation
	err       error
}

func (s stubRegistry) Get(name string) (calculator.Operation, error) {
	if s.err != nil {
		return nil, s.err
	}

	return s.operation, nil
}

func TestServiceExecute(t *testing.T) {
	service := NewService(stubRegistry{operation: calculator.AddOperation{}})

	response, err := service.Execute(Request{
		Operation: "add",
		Operands:  []float64{10, 5},
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if response.Result != 15 {
		t.Fatalf("expected result 15, got %v", response.Result)
	}
}

func TestServiceExecuteValidationError(t *testing.T) {
	service := NewService(stubRegistry{operation: calculator.AddOperation{}})

	_, err := service.Execute(Request{
		Operation: "add",
		Operands:  []float64{},
	})
	if !errors.Is(err, ErrInvalidRequest) {
		t.Fatalf("expected invalid request error, got %v", err)
	}
}
