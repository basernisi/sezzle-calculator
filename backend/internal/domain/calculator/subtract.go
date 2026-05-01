package calculator

type SubtractOperation struct{}

func (SubtractOperation) Name() string {
	return "subtract"
}

func (SubtractOperation) Validate(operands []float64) error {
	return validateExactlyTwoOperands(operands)
}

func (SubtractOperation) Execute(operands []float64) (float64, error) {
	return operands[0] - operands[1], nil
}
