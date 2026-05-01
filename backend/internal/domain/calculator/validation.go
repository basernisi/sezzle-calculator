package calculator

func validateExactlyOneOperand(operands []float64) error {
	if len(operands) != 1 {
		return ErrInvalidOperandCount
	}

	return nil
}

func validateExactlyTwoOperands(operands []float64) error {
	if len(operands) != 2 {
		return ErrInvalidOperandCount
	}

	return nil
}
