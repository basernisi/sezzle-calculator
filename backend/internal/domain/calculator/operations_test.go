package calculator

import (
	"errors"
	"testing"
)

func TestOperationsExecute(t *testing.T) {
	tests := []struct {
		name      string
		operation Operation
		operands  []float64
		expected  float64
		wantErr   error
	}{
		{name: "add", operation: AddOperation{}, operands: []float64{10, 5}, expected: 15},
		{name: "subtract", operation: SubtractOperation{}, operands: []float64{10, 5}, expected: 5},
		{name: "multiply", operation: MultiplyOperation{}, operands: []float64{10, 5}, expected: 50},
		{name: "divide", operation: DivideOperation{}, operands: []float64{10, 5}, expected: 2},
		{name: "power", operation: PowerOperation{}, operands: []float64{2, 3}, expected: 8},
		{name: "sqrt", operation: SquareRootOperation{}, operands: []float64{9}, expected: 3},
		{name: "percentage", operation: PercentageOperation{}, operands: []float64{20, 50}, expected: 10},
		{name: "division by zero", operation: DivideOperation{}, operands: []float64{10, 0}, wantErr: ErrDivisionByZero},
		{name: "negative sqrt", operation: SquareRootOperation{}, operands: []float64{-1}, wantErr: ErrNegativeSquareRoot},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.operation.Validate(test.operands); err != nil {
				t.Fatalf("expected validation to pass, got %v", err)
			}

			result, err := test.operation.Execute(test.operands)
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("expected error %v, got %v", test.wantErr, err)
			}

			if err == nil && result != test.expected {
				t.Fatalf("expected result %v, got %v", test.expected, result)
			}
		})
	}
}

func TestOperationValidateOperandCount(t *testing.T) {
	tests := []struct {
		name      string
		operation Operation
		operands  []float64
	}{
		{name: "add", operation: AddOperation{}, operands: []float64{1}},
		{name: "subtract", operation: SubtractOperation{}, operands: []float64{1}},
		{name: "multiply", operation: MultiplyOperation{}, operands: []float64{1}},
		{name: "divide", operation: DivideOperation{}, operands: []float64{1}},
		{name: "power", operation: PowerOperation{}, operands: []float64{1}},
		{name: "percentage", operation: PercentageOperation{}, operands: []float64{1}},
		{name: "sqrt", operation: SquareRootOperation{}, operands: []float64{1, 2}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.operation.Validate(test.operands)
			if !errors.Is(err, ErrInvalidOperandCount) {
				t.Fatalf("expected operand count error, got %v", err)
			}
		})
	}
}
