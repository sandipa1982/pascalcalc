package expression

import (
	"strings"
	"testing"
)

func TestExpressionLengthBoundary(t *testing.T) {
	e := NewEvaluator()
	// A valid 1024-byte expression consisting of a 1 and 1023 spaces.
	valid := "1" + strings.Repeat(" ", 1023)
	result, err := e.Evaluate(valid)
	if err != nil || result != 1 {
		t.Fatalf("1024-byte expression: got %v, %v; want 1", result, err)
	}
	oversized := valid + " "
	if _, err := e.Evaluate(oversized); err == nil || err.Error() != "expression too long" {
		t.Fatalf("1025-byte expression: got error %v; want expression too long", err)
	}
}
