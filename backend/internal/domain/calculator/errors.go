package calculator

import "errors"

var (
	ErrInvalidOperandCount  = errors.New("invalid operand count")
	ErrDivisionByZero       = errors.New("division by zero")
	ErrNegativeSquareRoot   = errors.New("square root of a negative number")
	ErrUnsupportedOperation = errors.New("unsupported operation")
)
