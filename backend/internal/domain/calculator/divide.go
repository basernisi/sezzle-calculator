package calculator

type DivideOperation struct{}

func (DivideOperation) Name() string {
	return "divide"
}

func (DivideOperation) Validate(operands []float64) error {
	return validateExactlyTwoOperands(operands)
}

func (DivideOperation) Execute(operands []float64) (float64, error) {
	if operands[1] == 0 {
		return 0, ErrDivisionByZero
	}

	return operands[0] / operands[1], nil
}
