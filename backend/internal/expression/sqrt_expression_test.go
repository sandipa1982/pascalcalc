package expression

import (
	"math"
	"testing"
)

func TestSqrtFunctionExpressions(t *testing.T) {
	cases := []struct {
		input string
		want  float64
	}{
		{"sqrt(16)", 4}, {"2+4+sqrt(16)", 10}, {"sqrt(9)+sqrt(16)", 7},
		{"2*sqrt(25)", 10}, {"sqrt(16+9)", 5}, {"sqrt(sqrt(16))", 2}, {"√(16)", 4},
	}
	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := NewEvaluator().Evaluate(tc.input)
			if err != nil || math.Abs(got-tc.want) > 1e-9 {
				t.Fatalf("Evaluate(%q)=%v, %v; want %v", tc.input, got, err, tc.want)
			}
		})
	}
}
func TestNegativeSquareRootExpression(t *testing.T) {
	if _, err := NewEvaluator().Evaluate("sqrt(-4)"); err == nil {
		t.Fatal("expected error")
	}
}
