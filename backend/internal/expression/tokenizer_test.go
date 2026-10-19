package expression

import "testing"

func TestTokenize(t *testing.T) {
	tokens, err := Tokenize("20 + 5 * 3 - 1")
	if err != nil {
		t.Fatal(err)
	}

	expected := []TokenType{
		TokenNumber,   // 20
		TokenPlus,     // +
		TokenNumber,   // 5
		TokenMultiply, // *
		TokenNumber,   // 3
		TokenMinus,    // -
		TokenNumber,   // 1
		TokenEOF,
	}

	if len(tokens) != len(expected) {
		t.Fatalf(
			"expected %d tokens, got %d",
			len(expected), len(tokens),
		)
	}

	for i, want := range expected {
		if tokens[i].Type != want {
			t.Errorf(
				"token %d: expected %v, got %v",
				i, want, tokens[i].Type,
			)
		}
	}
}
