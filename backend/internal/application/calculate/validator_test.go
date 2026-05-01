package calculate

import (
	"errors"
	"math"
	"testing"
)

func TestValidateRequest(t *testing.T) {
	tests := []struct {
		name    string
		request Request
		wantErr error
	}{
		{
			name: "valid request",
			request: Request{
				Operation: "add",
				Operands:  []float64{1, 2},
			},
		},
		{
			name: "missing operation",
			request: Request{
				Operands: []float64{1, 2},
			},
			wantErr: ErrInvalidRequest,
		},
		{
			name: "missing operands",
			request: Request{
				Operation: "add",
			},
			wantErr: ErrInvalidRequest,
		},
		{
			name: "nan operand",
			request: Request{
				Operation: "add",
				Operands:  []float64{math.NaN()},
			},
			wantErr: ErrInvalidRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateRequest(test.request)
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("expected error %v, got %v", test.wantErr, err)
			}
		})
	}
}
