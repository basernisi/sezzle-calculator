package calculator

import "math"

type SquareRootOperation struct{}

func (SquareRootOperation) Name() string {
	return "sqrt"
}

func (SquareRootOperation) Validate(operands []float64) error {
	return validateExactlyOneOperand(operands)
}

func (SquareRootOperation) Execute(operands []float64) (float64, error) {
	if operands[0] < 0 {
		return 0, ErrNegativeSquareRoot
	}

	return math.Sqrt(operands[0]), nil
}
