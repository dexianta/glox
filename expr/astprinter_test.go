package expr

import (
	"dexianta/glox/scanner"
	"fmt"
	"testing"
)

func TestAstPrinter(t *testing.T) {
	expr := Binary{
		Left:     Unary{
			Operator: scanner.Token{
				Type:    scanner.MINUS,
				Lexeme:  "-",
				Literal: nil,
				Line:    1,
			},
			Right:    Literal{Value: 1231},
		},
		Operator: scanner.Token{
			Type:    scanner.STAR,
			Lexeme:  "*",
			Literal: nil,
			Line:    1,
		},
		Right:    Grouping{
			Expression: Literal{Value: 42.12},
		},
	}

	fmt.Println(AstPrinter{}.Print(expr))
}
