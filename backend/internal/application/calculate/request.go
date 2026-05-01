package calculate

type Request struct {
	Operation string
	Operands  []float64
}

type Response struct {
	Result float64
}
