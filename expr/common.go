package expr

import (
	"fmt"
	"strconv"
)

type Expr interface{
	isExpr()
	accept(visitor Visitor) Value
}

type Value struct {
	object interface{}
}

func (v Value) string() string {
	return fmt.Sprintf("%v", v.object)
}

func (v Value) number() float64 {
	num, err := strconv.ParseFloat(fmt.Sprintf("%v", v.object), 64)
	if err != nil {
		panic(err.Error())
	}
	return num
}
