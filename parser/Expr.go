package parser

import (
	"dexianta/glox/scanner"
)

type Expr interface {
	isExpr()
	Accept(visitor Visitor) interface{}
}
//
//type Value struct {
//	object interface{}
//}
//
//func (v Value) string() string {
//	return fmt.Sprintf("%v", v.object)
//}
//
//func (v Value) number() float64 {
//	num, errorhandle := strconv.ParseFloat(fmt.Sprintf("%v", v.object), 64)
//	if errorhandle != nil {
//		panic(errorhandle.Error())
//	}
//	return num
//}

// ================================================ //
//						Visitors
// ================================================ //

type Visitor interface {
	VisitBinaryExpr(binary Binary) interface{}
	VisitGroupingExpr(grouping Grouping) interface{}
	VisitLiteralExpr(literal Literal) interface{}
	VisitUnaryExpr(u Unary) interface{}
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
func (b Binary) Accept(visitor Visitor) interface{}{
	return visitor.VisitBinaryExpr(b)
}

// ========================= //

type Grouping struct {
	Expression Expr
}

func (g Grouping) isExpr() {}
func (g Grouping) Accept(visitor Visitor) interface{} {
	return visitor.VisitGroupingExpr(g)
}

// ========================= //

type Literal struct {
	Value interface{}
}

func (l Literal) isExpr() {}
func (l Literal) Accept(visitor Visitor) interface{} {
	return visitor.VisitLiteralExpr(l)
}

// ========================= //

type Unary struct {
	Operator scanner.Token
	Right    Expr
}

func (u Unary) isExpr() {}
func (u Unary) Accept(visitor Visitor) interface{} {
	return visitor.VisitUnaryExpr(u)
}

//func logErr(line int, msg string) {
//	errorhandle.report(line, "", msg)
//}
