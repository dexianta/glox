package scanner

import (
	"github.com/stretchr/testify/assert"
	"testing"
)



func TestScanner(t *testing.T) {
	t.Run("scan brackets", func(t *testing.T) {
		scanner := NewScanner("(){}")
		tokens := scanner.ScanTokens()
		assert.Equal(t, tokens, []Token{
			{
				Type:   LEFT_PAREN,
				Lexeme: "(",
				Line:   0,
			},
			{
				Type:   RIGHT_PAREN,
				Lexeme: ")",
				Line:   0,
			},
			{
				Type:   LEFT_BRACE,
				Lexeme: "{",
				Line:   0,
			},
			{
				Type:   RIGHT_BRACE,
				Lexeme: "}",
				Line:   0,
			},
			{
				Type: EOF,
				Line: 0,
			},
		})
	})

	t.Run("scan brackets with comments", func(t *testing.T) {
		scanner := NewScanner("()//")
		tokens := scanner.ScanTokens()
		assert.Equal(t, tokens, []Token{
			{
				Type:   LEFT_PAREN,
				Lexeme: "(",
				Line:   0,
			},
			{
				Type:   RIGHT_PAREN,
				Lexeme: ")",
				Line:   0,
			},
			{
				Type: EOF,
				Line: 0,
			},
		})
	})

	t.Run("scan brackets with comments", func(t *testing.T) {
		scanner := NewScanner("()//()()()")
		tokens := scanner.ScanTokens()
		assert.Equal(t, tokens, []Token{
			{
				Type:   LEFT_PAREN,
				Lexeme: "(",
				Line:   0,
			},
			{
				Type:   RIGHT_PAREN,
				Lexeme: ")",
				Line:   0,
			},
			{
				Type: EOF,
				Line: 0,
			},
		})
	})

	t.Run("testing comments", func(t *testing.T) {
		scanner := NewScanner("()//()()()\n()")
		tokens := scanner.ScanTokens()
		assert.Equal(t, tokens, []Token{
			{
				Type:   LEFT_PAREN,
				Lexeme: "(",
				Line:   0,
			},
			{
				Type:   RIGHT_PAREN,
				Lexeme: ")",
				Line:   0,
			},
			{
				Type:   LEFT_PAREN,
				Lexeme: "(",
				Line:   1,
			},
			{
				Type:   RIGHT_PAREN,
				Lexeme: ")",
				Line:   1,
			},
			{
				Type: EOF,
				Line: 1,
			},
		})
	})

	t.Run("operators", func(t *testing.T) {
		scanner := NewScanner("+-/>=<=")
		tokens := scanner.ScanTokens()

		expectedTokens := []Token{
			{
				Type:    PLUS,
				Lexeme:  "+",
				Line:    0,
			},
			{
				Type:    MINUS,
				Lexeme:  "-",
				Line:    0,
			},
			{
				Type:    SLASH,
				Lexeme:  "/",
				Line:    0,
			},
			{
				Type:    GREATER_EQUAL,
				Lexeme:  ">=",
				Line:    0,
			},
			{
				Type:    LESS_EUQAL,
				Lexeme:  "<=",
				Line:    0,
			},
			{
				Type: EOF,
				Line: 0,
			},
		}

		assert.Equal(t, expectedTokens, tokens)
	})
}
