package arithmetic

import "errors"

// Service contains the arithmetic business logic.
type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Add(a, b float64) float64 {
	return a + b
}

func (s *Service) Subtract(a, b float64) float64 {
	return a - b
}

func (s *Service) Multiply(a, b float64) float64 {
	return a * b
}

func (s *Service) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}
