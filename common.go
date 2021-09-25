package main

import (
	"fmt"
	"strconv"
)

type Expr interface {
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

// ================================================ //
//						Visitors
// ================================================ //

type Visitor interface {
	visitBinary(binary Binary) Value
	visitGrouping(grouping Grouping) Value
	visitLiteral(literal Literal) Value
	visitUnary(u Unary) Value
}

// ========================= //

type Binary struct {
	Left     Expr
	Operator Token
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
	Operator Token
	Right    Expr
}

func (u Unary) isExpr() {}
func (u Unary) accept(visitor Visitor) Value {
	return visitor.visitUnary(u)
}

func logErr(line int, msg string) {
	report(line, "", msg)
}
