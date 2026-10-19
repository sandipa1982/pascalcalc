package mathops

import (
	"errors"
	"math"
)

// Service contains mathematical operations.
type Service struct{}

// NewService creates a Math Service instance.
func NewService() *Service {
	return &Service{}
}

// Power calculates a raised to b.
func (s *Service) Power(a, b float64) (float64, error) {
	if !isFinite(a) || !isFinite(b) {
		return 0, errors.New("inputs must be finite numbers")
	}

	result := math.Pow(a, b)

	if !isFinite(result) {
		return 0, errors.New("invalid exponentiation result")
	}

	return result, nil
}

// SquareRoot calculates the square root of a.
func (s *Service) SquareRoot(a float64) (float64, error) {
	if !isFinite(a) {
		return 0, errors.New("input must be a finite number")
	}

	if a < 0 {
		return 0, errors.New(
			"cannot calculate square root of a negative number",
		)
	}

	return math.Sqrt(a), nil
}

// Percentage calculates a percent of b.
// Example: Percentage(20, 150) = 30.
func (s *Service) Percentage(a, b float64) (float64, error) {
	if !isFinite(a) || !isFinite(b) {
		return 0, errors.New("inputs must be finite numbers")
	}

	result := (a / 100) * b

	if !isFinite(result) {
		return 0, errors.New("invalid percentage result")
	}

	return result, nil
}

// isFinite rejects NaN and positive/negative infinity.
func isFinite(value float64) bool {
	return !math.IsNaN(value) && !math.IsInf(value, 0)
}
