package expression

import (
	"fmt"
	"strconv"
	"unicode"
)

type TokenType int

const (
	TokenNumber TokenType = iota
	TokenPlus
	TokenMinus
	TokenMultiply
	TokenDivide
	TokenPower
	TokenLeftParen
	TokenRightParen
	TokenPercent
	TokenSqrt
	TokenEOF
)

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(input string) ([]Token, error) {
	var tokens []Token
	chars := []rune(input)

	for i := 0; i < len(chars); {
		ch := chars[i]
		if i+4 <= len(chars) && string(chars[i:i+4]) == "sqrt" {
			tokens = append(tokens, Token{Type: TokenSqrt, Value: "sqrt"})
			i += 4
			continue
		}

		if unicode.IsSpace(ch) {
			i++
			continue
		}

		if ch >= '0' && ch <= '9' || ch == '.' {
			start := i
			dots := 0

			for i < len(chars) {
				c := chars[i]
				if c >= '0' && c <= '9' {
					i++
				} else if c == '.' {
					dots++
					if dots > 1 {
						return nil, fmt.Errorf("invalid number at position %d", start)
					}
					i++
				} else {
					break
				}
			}

			value := string(chars[start:i])
			if _, err := strconv.ParseFloat(value, 64); err != nil {
				return nil, fmt.Errorf("invalid number %q: %w", value, err)
			}

			tokens = append(tokens, Token{
				Type: TokenNumber, Value: value,
			})
			continue
		}

		var tokenType TokenType

		switch ch {
		case '+':
			tokenType = TokenPlus
		case '-', '−':
			tokenType = TokenMinus
		case '*', '×':
			tokenType = TokenMultiply
		case '/', '÷':
			tokenType = TokenDivide
		case '^':
			tokenType = TokenPower
		case '(':
			tokenType = TokenLeftParen
		case ')':
			tokenType = TokenRightParen
		case '%':
			tokenType = TokenPercent
		case '√':
			tokenType = TokenSqrt
		default:
			return nil, fmt.Errorf(
				"unsupported character %q at position %d", ch, i,
			)
		}

		tokens = append(tokens, Token{
			Type: tokenType, Value: string(ch),
		})
		i++
	}

	tokens = append(tokens, Token{Type: TokenEOF})
	return tokens, nil
}
