package parser

import (
	"dexianta/glox/scanner"
)

type Expr interface {
	isExpr()
}
// ========================= //
// 			expression
// ========================= //

type Binary struct {
	Left     Expr
	Operator scanner.Token
	Right    Expr
}

func (b Binary) isExpr() {}

// ========================= //

type Grouping struct {
	Expression Expr
}

func (g Grouping) isExpr() {}

// ========================= //

type Literal struct {
	Value interface{}
}

func (l Literal) isExpr() {}

// ========================= //

type Unary struct {
	Operator scanner.Token
	Right    Expr
}

func (u Unary) isExpr() {}