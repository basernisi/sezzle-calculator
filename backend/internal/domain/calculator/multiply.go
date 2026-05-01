package calculator

type MultiplyOperation struct{}

func (MultiplyOperation) Name() string {
	return "multiply"
}

func (MultiplyOperation) Validate(operands []float64) error {
	return validateExactlyTwoOperands(operands)
}

func (MultiplyOperation) Execute(operands []float64) (float64, error) {
	return operands[0] * operands[1], nil
}
