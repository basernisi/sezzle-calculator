package calculator

type AddOperation struct{}

func (AddOperation) Name() string {
	return "add"
}

func (AddOperation) Validate(operands []float64) error {
	return validateExactlyTwoOperands(operands)
}

func (AddOperation) Execute(operands []float64) (float64, error) {
	return operands[0] + operands[1], nil
}
