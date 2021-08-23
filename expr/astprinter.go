package expr

import "strings"

type AstPrinter struct {}

func (a AstPrinter) Print(expr Expr) string {
	return expr.accept(a).string()
}

func (a AstPrinter) visitBinary(binary Binary) Value {
	return a.parenthesize(binary.Operator.Lexeme, binary.Left, binary.Right)
}

func (a AstPrinter) visitGrouping(grouping Grouping) Value {
	return a.parenthesize("group", grouping.Expression)
}

func (a AstPrinter) visitLiteral(literal Literal) Value {
	if literal.Value == nil {
		return Value{"nil"}
	}
	return Value{literal.Value}
}

func (a AstPrinter) visitUnary(u Unary) Value {
	return a.parenthesize(u.Operator.Lexeme, u.Right)
}

func (a AstPrinter) parenthesize(name string, exprs ...Expr) Value {
	sb := strings.Builder{}
	sb.WriteString("(")
	sb.WriteString(name)
	for _, expr := range exprs {
		sb.WriteString(" ")
		sb.WriteString(expr.accept(a).string())
	}
	sb.WriteString(")")
	return Value{object: sb.String()}
}


