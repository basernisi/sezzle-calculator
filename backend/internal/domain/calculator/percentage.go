package calculator

type PercentageOperation struct{}

func (PercentageOperation) Name() string {
	return "percentage"
}

func (PercentageOperation) Validate(operands []float64) error {
	return validateExactlyTwoOperands(operands)
}

func (PercentageOperation) Execute(operands []float64) (float64, error) {
	return (operands[0] / 100) * operands[1], nil
}
