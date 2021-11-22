package interpreter

import (
    "dexianta/glox/parser"
    "dexianta/glox/scanner"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestBinaryExpr(t *testing.T) {
    expr := parser.Binary{
        Left:     parser.Literal{Value: float64(3)},
        Operator: scanner.Token{
            Type:    scanner.PLUS,
            Lexeme:  "+",
            Literal: nil,
            Line:    0,
        },
        Right:    parser.Literal{Value: float64(5)},
    }

    res, err := BinaryExpr(expr)
    assert.Nil(t, err)
    assert.Equal(t, res, float64(8))
}
