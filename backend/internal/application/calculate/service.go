package calculate

import "github.com/jnsilvag/sezzle-calculator/backend/internal/ports"

type Service struct {
	registry ports.OperationRegistry
}

func NewService(registry ports.OperationRegistry) Service {
	return Service{registry: registry}
}

func (s Service) Execute(request Request) (Response, error) {
	if err := ValidateRequest(request); err != nil {
		return Response{}, err
	}

	operation, err := s.registry.Get(request.Operation)
	if err != nil {
		return Response{}, err
	}

	if err := operation.Validate(request.Operands); err != nil {
		return Response{}, err
	}

	result, err := operation.Execute(request.Operands)
	if err != nil {
		return Response{}, err
	}

	return Response{Result: result}, nil
}
