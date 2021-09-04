package main

import (
	"fmt"
	"testing"
)

func TestAstPrinter(t *testing.T) {
	expr := Binary{
		Left:     Unary{
			Operator: main.Token{
				Type:    main.MINUS,
				Lexeme:  "-",
				Literal: nil,
				Line:    1,
			},
			Right: Literal{Value: 1231},
		},
		Operator: main.Token{
			Type:    main.STAR,
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
