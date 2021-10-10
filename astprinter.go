package main

import (
	"dexianta/glox/parser"
	"strings"
)

type AstPrinter struct{}

func (a AstPrinter) Print(expr parser.Expr) string {
	return expr.Accept(a).(string)
}

func (a AstPrinter) VisitBinaryExpr(binary parser.Binary) interface{} {
	return a.parenthesize(binary.Operator.Lexeme, binary.Left, binary.Right)
}

func (a AstPrinter) VisitGroupingExpr(grouping parser.Grouping) interface{} {
	return a.parenthesize("group", grouping.Expression)
}

func (a AstPrinter) VisitLiteralExpr(literal parser.Literal) interface{} {
	if literal.Value == nil {
		return "nil"
	}
	return literal.Value
}

func (a AstPrinter) VisitUnaryExpr(u parser.Unary) interface{} {
	return a.parenthesize(u.Operator.Lexeme, u.Right)
}

func (a AstPrinter) parenthesize(name string, exprs ...parser.Expr) interface{} {
	sb := strings.Builder{}
	sb.WriteString("(")
	sb.WriteString(name)
	for _, expr := range exprs {
		sb.WriteString(" ")
		sb.WriteString(expr.Accept(a).(string))
	}
	sb.WriteString(")")
	return sb.String()
}
