package parser

import (
	"dexianta/glox/errorhandle"
	"dexianta/glox/scanner"
	"errors"
)

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
	tokens  []scanner.Token
}

func NewParser(tokens []scanner.Token) Parser {
	return Parser{
		tokens: tokens,
	}
}

func (p *Parser) Parse() Expr {
	expr, err := p.expr()
	switch err {
	case ParseError:
		return nil
	default:
		return expr
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

	for p.match(scanner.BANG_EQUAL, scanner.EQUAL_EQUAL) {
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

	for p.match(scanner.GREATER, scanner.GREATER_EQUAL, scanner.LESS, scanner.LESS_EQUAL) {
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

	for p.match(scanner.MINUS, scanner.PLUS) {
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

	for p.match(scanner.SLASH, scanner.STAR) {
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
	if p.match(scanner.BANG, scanner.MINUS) {
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
	if p.match(scanner.FALSE) {
		return Literal{false}, nil
	}
	if p.match(scanner.TRUE) {
		return Literal{true}, nil
	}
	if p.match(scanner.NIL) {
		return Literal{nil}, nil
	}

	if p.match(scanner.NUMBER, scanner.STRING) {
		return Literal{p.previous().Literal}, nil
	}

	if p.match(scanner.LEFT_PAREN) {
		expr, err := p.expr()
		if err != nil {
			return expr, err
		}
		p.consume(scanner.RIGHT_PAREN, "Expect ')' after expression")
		return Grouping{expr}, nil
	}

	return nil, p.error(p.peek(), "expect expression")
}

// ===========================================
// helpers
// ===========================================

var ParseError = errors.New("parse error")
func (p *Parser) consume(tokenType scanner.TokenType, msg string) (scanner.Token, error) {
	if p.check(tokenType) {
		return p.advance(), nil
	}

	return scanner.Token{}, p.error(p.peek(), msg)
}

func (p *Parser) error(token scanner.Token, msg string) error {
	if token.Type == scanner.EOF {
		errorhandle.Report(token.Line, " at end", msg)
	} else {
		errorhandle.Report(token.Line, "at '"+token.Lexeme+"'", msg)
	}
	return ParseError
}

func (p *Parser) match(types ...scanner.TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(t scanner.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == t
}

func (p *Parser) peek() scanner.Token {
	return p.tokens[p.current]
}

func (p *Parser) advance() scanner.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == scanner.EOF
}

func (p *Parser) previous() scanner.Token {
	return p.tokens[p.current-1]
}

// keep discarding the token until we hit the beginning of the next statement
func (p *Parser) sync() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().Type == scanner.SEMICOLON {
			return
		}

		switch p.peek().Type {
		case scanner.CLASS, scanner.FUN, scanner.VAR, scanner.FOR, scanner.IF, scanner.WHILE, scanner.PRINT, scanner.RETURN:
			return
		default:
			p.advance()
		}
	}
}
