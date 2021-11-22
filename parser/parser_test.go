package parser

import (
    "dexianta/glox/scanner"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestParser_Parse(t *testing.T) {
    // (1 + 2) * (3 - 5)
    tokens := []scanner.Token{
        {
            Type:    scanner.LEFT_PAREN,
            Lexeme:  "(",
        },
        {
            Type:    scanner.NUMBER,
            Lexeme:  "1",
            Literal: float64(1),
        },
        {
            Type:    scanner.PLUS,
            Lexeme:  "+",
        },
        {
            Type:    scanner.NUMBER,
            Lexeme:  "2",
            Literal: float64(2),
        },
        {
            Type:    scanner.RIGHT_PAREN,
            Lexeme:  ")",
        },
        {
            Type:    scanner.STAR,
            Lexeme:  "*",
        },
        {
            Type:    scanner.LEFT_PAREN,
            Lexeme:  "(",
        },
        {
            Type:    scanner.NUMBER,
            Lexeme:  "3",
            Literal: float64(3),
        },
        {
            Type:    scanner.MINUS,
            Lexeme:  "-",
        },
        {
            Type:    scanner.NUMBER,
            Lexeme:  "5",
            Literal: float64(5),
        },
        {
            Type:    scanner.RIGHT_PAREN,
            Lexeme:  ")",
        },
        {
            Type: scanner.EOF,
        },
    }
    parser := NewParser(tokens)
    expr := parser.Parse()

    expected := Binary{
        Left:
            Grouping{Binary{
            Left:     Literal{Value: float64(1)},
            Operator: scanner.Token{
                Type:    scanner.PLUS,
                Lexeme:  "+",
            },
            Right:    Literal{Value: float64(2)},
        }},
        Operator: scanner.Token{
            Type:    scanner.STAR,
            Lexeme:  "*",
        },
        Right: Grouping{Binary{
            Left:     Literal{Value: float64(3)},
            Operator: scanner.Token{
                Type:    scanner.MINUS,
                Lexeme:  "-",
            },
            Right:    Literal{Value: float64(5)},
        }},
    }
    assert.Equal(t, expected, expr)
}
