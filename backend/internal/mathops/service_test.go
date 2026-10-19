package mathops

import (
	"math"
	"testing"
)

func TestPower(t *testing.T) {
	service := NewService()

	result, err := service.Power(2, 3)

	if err != nil || result != 8 {
		t.Fatalf("expected 8, got %v, error: %v", result, err)
	}
}

func TestPowerInvalidResult(t *testing.T) {
	service := NewService()

	_, err := service.Power(10, 1000)

	if err == nil {
		t.Fatal("expected error for exponentiation overflow")
	}
}

func TestSquareRoot(t *testing.T) {
	service := NewService()

	result, err := service.SquareRoot(81)

	if err != nil || result != 9 {
		t.Fatalf("expected 9, got %v, error: %v", result, err)
	}
}

func TestSquareRootNegative(t *testing.T) {
	service := NewService()

	_, err := service.SquareRoot(-9)

	if err == nil {
		t.Fatal("expected error for negative square root")
	}
}

func TestPercentage(t *testing.T) {
	service := NewService()

	result, err := service.Percentage(20, 150)

	if err != nil || math.Abs(result-30) > 1e-9 {
		t.Fatalf("expected 30, got %v, error: %v", result, err)
	}
}

func TestPercentageInvalidInput(t *testing.T) {
	service := NewService()

	_, err := service.Percentage(math.NaN(), 100)

	if err == nil {
		t.Fatal("expected error for NaN input")
	}
}
