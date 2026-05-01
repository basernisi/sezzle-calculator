package calculator

type Operation interface {
	Name() string
	Validate(operands []float64) error
	Execute(operands []float64) (float64, error)
}
