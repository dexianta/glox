package interpreter

import (
	"dexianta/glox/parser"
	"dexianta/glox/scanner"
	"fmt"
	"reflect"
)

func Interpret(expr parser.Expr) (res interface{}, err error) {
	switch expr.(type) {
	case parser.Binary:
		res, err = BinaryExpr(expr.(parser.Binary))
	case parser.Unary:
		res, err = UnaryExpr(expr.(parser.Unary))
	case parser.Grouping:
		res, err = GroupingExpr(expr.(parser.Grouping))
	case parser.Literal:
		res, err = LiteralExpr(expr.(parser.Literal))
	default:
		err = RuntimeError{Msg: "invalid expr"}
	}

	if err != nil {
		fmt.Println(err)
	}
	return res, err
}

func BinaryExpr(binary parser.Binary) (interface{}, error) {
	left, err := Interpret(binary.Left)
	if err != nil {
		return nil, err
	}
	right, err := Interpret(binary.Right)
	if err != nil {
		return nil, err
	}

	op := binary.Operator
	switch op.Type {
	case scanner.MINUS:
		err := checkNumberOperands(op, right, left)
		if err != nil {
			return nil, err
		}
		return left.(float64) - right.(float64), nil
	case scanner.SLASH:
		err := checkNumberOperands(op, right, left)
		if err != nil {
			return nil, err
		}
		return left.(float64) / right.(float64), nil
	case scanner.STAR:
		err := checkNumberOperands(op, right, left)
		if err != nil {
			return nil, err
		}
		return left.(float64) * right.(float64), nil
	case scanner.PLUS:
		n1, ok1 := left.(float64)
		n2, ok2 := right.(float64)
		if ok1 && ok2 {
			return n1 + n2, nil
		}

		s1, ok1 := left.(string)
		s2, ok2 := right.(string)
		if ok1 && ok2 {
			return s1 + s2, nil
		}
	case scanner.GREATER:
		err := checkNumberOperands(op, right, left)
		if err != nil {
			return nil, err
		}
		return left.(float64) > right.(float64), nil
	case scanner.GREATER_EQUAL:
		err := checkNumberOperands(op, right, left)
		if err != nil {
			return nil, err
		}
		return left.(float64) >= right.(float64), nil
	case scanner.LESS:
		err := checkNumberOperands(op, right, left)
		if err != nil {
			return nil, err
		}
		return left.(float64) < right.(float64), nil
	case scanner.LESS_EQUAL:
		err := checkNumberOperands(op, right, left)
		if err != nil {
			return nil, err
		}
		return left.(float64) <= right.(float64), nil
	case scanner.BANG_EQUAL:
		return !isEqual(left, right), nil
	case scanner.EQUAL_EQUAL:
		return isEqual(left, right), nil
	default:
		return nil, RuntimeError{
			Token: op,
			Msg:   "didn't match any operator",
		}
	}
	return nil, RuntimeError{
		Token: op,
		Msg:   "not sure what's wrong",
	}
}

func GroupingExpr(grouping parser.Grouping) (interface{}, error) {
	return Interpret(grouping.Expression)
}

func LiteralExpr(literal parser.Literal) (interface{}, error) {
	return literal.Value, nil
}

func UnaryExpr(u parser.Unary) (interface{}, error) {
	right, err := Interpret(u.Right)
	if err != nil {
		return nil, err
	}
	switch u.Operator.Type {
	case scanner.MINUS:
		err := checkNumberOperand(u.Operator, right)
		if err != nil {
			return nil, err
		}
		return -(right.(float64)), nil
	case scanner.PLUS:
		err := checkNumberOperand(u.Operator, right)
		if err != nil {
			return nil, err
		}
		return right.(float64), nil
	case scanner.BANG:
		return isTruthy(right), nil
	default:
		return nil, RuntimeError{
			Token: u.Operator,
			Msg:   "invalid operator type",
		}
	}
}

func checkNumberOperands(operator scanner.Token, op1, op2 interface{}) error {
	switch op1.(type) {
	case float64:
		switch op2.(type) {
		case float64:
			return nil
		}
	}
	return RuntimeError{Token: operator, Msg: fmt.Sprintf("%v or %v is not a number", op1, op2)}
}

func checkNumberOperand(operator scanner.Token, num interface{}) error {
	switch num.(type) {
	case float64:
		return nil
	default:
		return RuntimeError{Token: operator, Msg: "%v is not a number"}
	}
}

func isEqual(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

func isTruthy(o interface{}) bool {
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