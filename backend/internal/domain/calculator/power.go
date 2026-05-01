package calculator

import "math"

type PowerOperation struct{}

func (PowerOperation) Name() string {
	return "power"
}

func (PowerOperation) Validate(operands []float64) error {
	return validateExactlyTwoOperands(operands)
}

func (PowerOperation) Execute(operands []float64) (float64, error) {
	return math.Pow(operands[0], operands[1]), nil
}
