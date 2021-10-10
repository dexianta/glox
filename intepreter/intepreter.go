package intepreter

import (
	"dexianta/glox/parser"
	"dexianta/glox/scanner"
)

type Interpreter struct {}

func (I *Interpreter) VisitBinaryExpr(binary parser.Binary) interface{} {
	left := I.evaluate(binary.Left)
	right := I.evaluate(binary.Right)

	switch binary.Operator.Type {
	case scanner.MINUS:
		return left.(float64) - right.(float64)
	case scanner.SLASH:
		return left.(float64) / right.(float64)
	case scanner.STAR:
		return left.(float64) * right.(float64)
	case scanner.PLUS:
		n1, ok1 := left.(float64)
		n2, ok2 := left.(float64)
		if ok1 && ok2 {
			return n1 + n2
		}

		s1, ok1 := left.(string)
		s2, ok2 := left.(string)
		if ok1 && ok2 {
			return s1 + s2
		}
	case scanner.GREATER:
		return left.(float64) > right.(float64)
	case scanner.GREATER_EQUAL:
		return left.(float64) >= right.(float64)
	case scanner.LESS:
		return left.(float64) < right.(float64)
	case scanner.LESS_EQUAL:
		return left.(float64) <= right.(float64)
	case scanner.BANG_EQUAL:
		return !I.isEqual(left, right)
	case scanner.EQUAL_EQUAL:
		return I.isEqual(left, right)
	default:
		return nil
	}
	return nil
}

func (I *Interpreter) VisitGroupingExpr(grouping parser.Grouping) interface{} {
	return I.evaluate(grouping.Expression)
}

func (I *Interpreter) VisitLiteralExpr(literal parser.Literal) interface{} {
	return literal.Value
}

func (I *Interpreter) VisitUnaryExpr(u parser.Unary) interface{} {
	right := I.evaluate(u.Right)
	switch u.Operator.Type {
	case scanner.MINUS:
		return -(right.(float64))
	case scanner.PLUS:
		return right.(float64)
	case scanner.BANG:
		return !I.isTruthy(right)
	default:
		return nil
	}
}

func (I *Interpreter) isEqual(a, b interface{}) bool {
	return a == b
}

func (I *Interpreter) isTruthy(o interface{}) bool {
	//TODO: see if can be put into the switch as well
	if o == nil {
		return false
	}
	switch o.(type) {
	case bool:
		return o.(bool)
	default:
		return true
	}
}

func (I *Interpreter) evaluate(e parser.Expr) interface{} {
	return e.Accept(I)
}

