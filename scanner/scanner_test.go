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

	t.Run("scan brackets with multiline comments", func(t *testing.T) {
		scanner := NewScanner("()/*()\n()\n()*/()")
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
				Line:   2,
			},
			{
				Type:   RIGHT_PAREN,
				Lexeme: ")",
				Line:   2,
			},
			{
				Type: EOF,
				Line: 2,
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
				Type:   PLUS,
				Lexeme: "+",
				Line:   0,
			},
			{
				Type:   MINUS,
				Lexeme: "-",
				Line:   0,
			},
			{
				Type:   SLASH,
				Lexeme: "/",
				Line:   0,
			},
			{
				Type:   GREATER_EQUAL,
				Lexeme: ">=",
				Line:   0,
			},
			{
				Type:   LESS_EQUAL,
				Lexeme: "<=",
				Line:   0,
			},
			{
				Type: EOF,
				Line: 0,
			},
		}

		assert.Equal(t, expectedTokens, tokens)
	})

	t.Run("test string", func(t *testing.T) {
		scanner := NewScanner("\"hello world\"\n//\"hello world\"")
		tokens := scanner.ScanTokens()

		expectedToken := []Token{{
			Type:    STRING,
			Lexeme:  "\"hello world\"",
			Literal: "hello world",
			Line:    0,
		},
			{
				Type: EOF,
				Line: 1,
			},
		}

		assert.Equal(t, tokens, expectedToken)
	})

	t.Run("test number", func(t *testing.T) {
		scanner := NewScanner("32")
		tokens := scanner.ScanTokens()

		expectedToken := []Token{{
			Type:    NUMBER,
			Lexeme:  "32",
			Literal: 32.0,
			Line:    0,
		},
			{
				Type: EOF,
				Line: 0,
			},
		}

		assert.Equal(t, tokens, expectedToken)
	})

	t.Run("test number, decimal number", func(t *testing.T) {
		scanner := NewScanner("32.123")
		tokens := scanner.ScanTokens()

		expectedToken := []Token{
			{
				Type:    NUMBER,
				Lexeme:  "32.123",
				Literal: 32.123,
				Line:    0,
			},
			{
				Type: EOF,
				Line: 0,
			},
		}

		assert.Equal(t, tokens, expectedToken)
	})

	t.Run("test number, multiple decimal number", func(t *testing.T) {
		scanner := NewScanner("32.123 546.123")
		tokens := scanner.ScanTokens()

		expectedToken := []Token{
			{
				Type:    NUMBER,
				Lexeme:  "32.123",
				Literal: 32.123,
				Line:    0,
			},

			{
				Type:    NUMBER,
				Lexeme:  "546.123",
				Literal: 546.123,
				Line:    0,
			},

			{
				Type: EOF,
				Line: 0,
			},
		}

		assert.Equal(t, tokens, expectedToken)
	})

	t.Run("identifier", func(t *testing.T) {
		scanner := NewScanner("if {hello} else {world}")
		tokens := scanner.ScanTokens()

		expectedToken := []Token{
			{
				Type:   IF,
				Lexeme: "if",
				Line:   0,
			},

			{
				Type:   LEFT_BRACE,
				Lexeme: "{",
				Line:   0,
			},

			{
				Type:   IDENTIFIER,
				Lexeme: "hello",
				Line:   0,
			},

			{
				Type:   RIGHT_BRACE,
				Lexeme: "}",
				Line:   0,
			},

			{
				Type:   ELSE,
				Lexeme: "else",
				Line:   0,
			},

			{
				Type:   LEFT_BRACE,
				Lexeme: "{",
				Line:   0,
			},

			{
				Type:   IDENTIFIER,
				Lexeme: "world",
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
		}

		assert.Equal(t, tokens, expectedToken)
	})
}
