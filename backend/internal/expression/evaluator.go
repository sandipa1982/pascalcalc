package expression

import (
	"calculator-app/backend/internal/arithmetic"
	"calculator-app/backend/internal/mathops"
	"errors"
	"fmt"
	"math"
	"strconv"
)

type Evaluator struct {
	tokens []Token
	pos    int

	arith *arithmetic.Service
	math  *mathops.Service

	maxExpressionLength int
}

const defaultMaxExpressionLength = 1024

func NewEvaluator() *Evaluator {
	return NewEvaluatorWithLimit(defaultMaxExpressionLength)
}

func NewEvaluatorWithLimit(maxLength int) *Evaluator {
	return &Evaluator{
		arith: arithmetic.NewService(),
		math:  mathops.NewService(),

		maxExpressionLength: maxLength,
	}
}
func (e *Evaluator) Evaluate(input string) (float64, error) {
	if len(input) > e.maxExpressionLength {
		return 0, errors.New("expression too long")
	}

	tokens, err := Tokenize(input)
	if err != nil {
		return 0, err
	}

	e.tokens = tokens
	e.pos = 0

	result, err := e.parseExpression()
	if err != nil {
		return 0, err
	}

	if e.current().Type != TokenEOF {
		return 0, fmt.Errorf(
			"unexpected token %q", e.current().Value,
		)
	}

	if math.IsNaN(result) || math.IsInf(result, 0) {
		return 0, errors.New("invalid calculation result")
	}

	return result, nil
}

func (e *Evaluator) current() Token {
	return e.tokens[e.pos]
}

func (e *Evaluator) advance() {
	if e.pos < len(e.tokens)-1 {
		e.pos++
	}
}

// Addition and subtraction (lowest precedence).
func (e *Evaluator) parseExpression() (float64, error) {
	left, err := e.parseTerm()
	if err != nil {
		return 0, err
	}

	for {
		op := e.current().Type
		if op != TokenPlus && op != TokenMinus {
			break
		}

		e.advance()
		right, err := e.parseTerm()
		if err != nil {
			return 0, err
		}

		if op == TokenPlus {
			left = e.arith.Add(left, right)
		} else {
			left = e.arith.Subtract(left, right)
		}
	}

	return left, nil
}

// Multiplication and division.
func (e *Evaluator) parseTerm() (float64, error) {
	left, err := e.parseUnary()
	if err != nil {
		return 0, err
	}

	for {
		op := e.current().Type
		if op != TokenMultiply && op != TokenDivide {
			break
		}

		e.advance()
		right, err := e.parseUnary()
		if err != nil {
			return 0, err
		}

		if op == TokenMultiply {
			left = e.arith.Multiply(left, right)
		} else {
			left, err = e.arith.Divide(left, right)
			if err != nil {
				return 0, err
			}
		}
	}

	return left, nil
}

// Unary minus and square root.
func (e *Evaluator) parseUnary() (float64, error) {
	switch e.current().Type {
	case TokenMinus:
		e.advance()
		value, err := e.parseUnary()
		return -value, err

	case TokenPlus:
		e.advance()
		return e.parseUnary()

	case TokenSqrt:
		e.advance()
		value, err := e.parseUnary()
		if err != nil {
			return 0, err
		}
		return e.math.SquareRoot(value)
	}

	return e.parsePower()
}

// Exponentiation.
func (e *Evaluator) parsePower() (float64, error) {
	left, err := e.parsePrimary()
	if err != nil {
		return 0, err
	}

	if e.current().Type == TokenPower {
		e.advance()
		right, err := e.parseUnary()
		if err != nil {
			return 0, err
		}
		return e.math.Power(left, right)
	}

	return left, nil
}

// Number, parentheses, and postfix percentage.
func (e *Evaluator) parsePrimary() (float64, error) {
	var value float64

	switch e.current().Type {
	case TokenNumber:
		parsed, err := strconv.ParseFloat(e.current().Value, 64)
		if err != nil {
			return 0, err
		}
		value = parsed
		e.advance()

	case TokenLeftParen:
		e.advance()
		parsed, err := e.parseExpression()
		if err != nil {
			return 0, err
		}
		if e.current().Type != TokenRightParen {
			return 0, errors.New("missing closing parenthesis")
		}
		e.advance()
		value = parsed

	default:
		return 0, fmt.Errorf(
			"expected number, got %q", e.current().Value,
		)
	}

	for e.current().Type == TokenPercent {
		value /= 100
		e.advance()
	}

	return value, nil
}
