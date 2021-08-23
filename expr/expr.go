package expr

import (
	"dexianta/glox/scanner"
)

type Visitor interface{
	visitBinary(binary Binary) Value
	visitGrouping(grouping Grouping) Value
	visitLiteral(literal Literal) Value
	visitUnary(u Unary) Value
}

// ========================= //
type Binary struct {
	Left     Expr
	Operator scanner.Token
	Right    Expr
}
func (b Binary) isExpr() {}
func (b Binary) accept(visitor Visitor) Value {
	return visitor.visitBinary(b)
}

// ========================= //
type Grouping struct {
	Expression Expr
}
func (g Grouping) isExpr() {}
func (g Grouping) accept(visitor Visitor) Value {
	return visitor.visitGrouping(g)
}

// ========================= //
type Literal struct {
	Value interface{}
}
func (l Literal) isExpr() {}
func (l Literal) accept(visitor Visitor) Value {
	return visitor.visitLiteral(l)
}

// ========================= //
type Unary struct {
	Operator scanner.Token
	Right Expr
}
func (u Unary) isExpr() {}
func (u Unary) accept(visitor Visitor) Value {
	return visitor.visitUnary(u)
}

