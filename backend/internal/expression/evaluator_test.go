package expression

import (
	"math"
	"testing"
)

func TestEvaluateExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"20 + 5 * 3 - 1", 34},
		{"10 / 2 + 3", 8},
		{"2 ^ 3 + 4", 12},
		{"√81 + 5", 14},
		{"200 * 20%", 40},
		{"5 * -2", -10},
		{"12.75 + 0.25", 13},
		{"2 ^ 3 ^ 2", 512},
		{"-2 ^ 2", -4},
	}

	evaluator := NewEvaluator()

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			result, err := evaluator.Evaluate(tc.input)
			if err != nil {
				t.Fatal(err)
			}

			if math.Abs(result-tc.expected) > 1e-9 {
				t.Fatalf(
					"expected %v, got %v",
					tc.expected, result,
				)
			}
		})
	}
}

func TestEvaluateErrors(t *testing.T) {
	tests := []string{
		"10 / 0",
		"√-9",
		"10 +",
		"2 @ 3",
		"12.3.4",
	}

	evaluator := NewEvaluator()

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			_, err := evaluator.Evaluate(input)
			if err == nil {
				t.Fatal("expected an error")
			}
		})
	}
}

func TestEvaluatorConfigurableLength(t *testing.T) {
	evaluator := NewEvaluatorWithLimit(5)

	_, err := evaluator.Evaluate("100 + 200")
	if err == nil {
		t.Fatal("expected expression length error")
	}

	evaluator = NewEvaluatorWithLimit(20)

	result, err := evaluator.Evaluate("100 + 200")
	if err != nil {
		t.Fatal(err)
	}

	if result != 300 {
		t.Fatalf("expected 300, got %v", result)
	}
}
