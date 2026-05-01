package calculate

import (
	"errors"
	"math"
	"strings"
)

var ErrInvalidRequest = errors.New("invalid request")

func ValidateRequest(request Request) error {
	if strings.TrimSpace(request.Operation) == "" {
		return ErrInvalidRequest
	}

	if len(request.Operands) == 0 {
		return ErrInvalidRequest
	}

	for _, operand := range request.Operands {
		if math.IsNaN(operand) || math.IsInf(operand, 0) {
			return ErrInvalidRequest
		}
	}

	return nil
}
