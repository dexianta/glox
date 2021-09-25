package main

import "errors"

// syntax tree
// ===========================================================
// expression     → equality ;
// equality       → comparison ( ( "!=" | "==" ) comparison )* ;
// comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
// term           → factor ( ( "-" | "+" ) factor )* ;
// factor         → unary ( ( "/" | "*" ) unary )* ;
// unary          → ( "!" | "-" ) unary | primary ;
// primary        → NUMBER | STRING | "true" | "false" | "nil" | "(" expression ")" ;

type Parser struct {
	current int
	tokens  []Token
}

func NewParser(tokens []Token) Parser {
	return Parser{
		tokens: tokens,
	}
}

func (p *Parser) Parse() (Expr, error) {
	expr, err := p.expr()
	switch err {
	case ParseError:
		return nil, err
	default:
		return expr, err
	}
}

func (p *Parser) expr() (Expr, error) {
	return p.equality()
}

func (p *Parser) equality() (Expr, error) {
	expr, err := p.comparison()
	if err != nil {
		return expr, err
	}

	for p.match(BANG_EQUAL, EQUAL_EQUAL) {
		operator := p.previous()
		right, err := p.comparison()
		if err != nil {
			return right, err
		}
		expr = Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr, nil
}

func (p *Parser) comparison() (Expr, error) {
	expr, err := p.term()
	if err != nil {
		return expr, err
	}

	for p.match(GREATER, GREATER_EQUAL, LESS, LESS_EUQAL) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return expr, err
		}
		expr = Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr, nil
}

func (p *Parser) term() (Expr, error) {
	expr, err := p.factor()
	if err != nil {
		return expr, err
	}

	for p.match(MINUS, PLUS) {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return right, err
		}
		expr = Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr, nil
}

func (p *Parser) factor() (Expr, error) {
	expr, err := p.unary()
	if err != nil {
		return expr, err
	}

	for p.match(SLASH, STAR) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return expr, err
		}
		expr = Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr, nil
}

func (p *Parser) unary() (Expr, error) {
	if p.match(BANG, MINUS) {
		operator := p.previous()
		right, err := p.unary()
		return Unary{
			Operator: operator,
			Right:    right,
		}, err
	}

	return p.primary()
}

func (p *Parser) primary() (Expr, error) {
	if p.match(FALSE) {
		return Literal{false}, nil
	}
	if p.match(TRUE) {
		return Literal{true}, nil
	}
	if p.match(NIL) {
		return Literal{nil}, nil
	}

	if p.match(NUMBER, STRING) {
		return Literal{p.previous().Literal}, nil
	}

	if p.match(LEFT_PAREN) {
		expr, err := p.expr()
		if err != nil {
			return expr, err
		}
		p.consume(RIGHT_PAREN, "Expect ')' after expression")
		return Grouping{expr}, nil
	}

	return nil, p.error(p.peek(), "expect expression")
}

func (p *Parser) consume(tokenType TokenType, msg string) (Token, error) {
	if p.check(tokenType) {
		return p.advance(), nil
	}

	return Token{}, p.error(p.peek(), msg)
}

func (p *Parser) error(token Token, msg string) error {
	LoxError(token, msg)
	return ParseError
}

var ParseError = errors.New("parse error")

// ===========================================
// helpers
// ===========================================

func (p *Parser) match(types ...TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(t TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == t
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == EOF
}

func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}

// keep discarding the token until we hit the beginning of the next statement
func (p *Parser) sync() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().Type == SEMICOLON {
			return
		}

		switch p.peek().Type {
		case CLASS, FUN, VAR, FOR, IF, WHILE, PRINT, RETURN:
			return
		default:
			p.advance()
		}
	}
}
