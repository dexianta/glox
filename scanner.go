package main

type TokenType uint

const (
	// single character tokens
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// one or two character tokens
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GERATER
	GREATER_EQUAL
	LESS
	LESS_EUQAL

	// literals
	INDETIFIER
	STRING
	NUMBER

	// keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
	EOF
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

type Scanner struct {
	Source  string
	Tokens  []Token
	start   int
	current int
	line    int
}

func NewScanner(source string) Scanner {
	return Scanner{Source: source}
}

func (s *Scanner) scanTokens() []Token {
	for !s.IsAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.Tokens = append(s.Tokens, Token{
		Type:    EOF,
		Lexeme:  "",
		Literal: nil,
		Line:    s.line})

	return s.Tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(': s.addToken(LEFT_PAREN, nil)
	case ')': s.addToken(RIGHT_PAREN, nil)
	case '{': s.addToken(LEFT_BRACE, nil)
	case '}': s.addToken(RIGHT_BRACE, nil)
	case ',': s.addToken(COMMA, nil)
	case '.': s.addToken(DOT, nil)
	case '-': s.addToken(MINUS, nil)
	case '+': s.addToken(PLUS, nil)
	case '*': s.addToken(STAR, nil)
	case ';': s.addToken(SEMICOLON, nil)

	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL, nil)
		}
		s.addToken(BANG, nil)

	case '/':
		if s.match('/') {
			for !s.match('\n') && !s.IsAtEnd() {
				s.advance()
			}
		} else if s.match('*') {
			for !s.match('*', '/') && !s.IsAtEnd(){
				s.advance()
			}
		} else {
			s.addToken(SLASH, nil)
		}

	default:
		logErr(s.line, "Unexpected character")

	}
}

func (s *Scanner) IsAtEnd() bool {
	return s.current >= len(s.Source)
}

func (s *Scanner) IsAtEndOffset(offset int) bool {
	return s.current + offset >= len(s.Source)
}

func (s *Scanner) advance() byte {
	idx := s.current
	s.current++
	return s.Source[idx]
}

func (s *Scanner) peek() (byte, bool) {
	if s.IsAtEnd() {
		return '0', true
	}
	return s.Source[s.current], false
}

//func(s *Scanner) match(char byte) bool {
//	if s.IsAtEnd() {
//		return false
//	}
//	if s.Source[s.current] != char {
//		return false
//	}
//
//	s.current++
//	return true
//}

func (s *Scanner) match(chars ...byte) bool {
	if s.IsAtEnd() && s.IsAtEndOffset(len(chars)){
		return false
	}

	for idx, c := range chars {
		if s.Source[s.current + idx] != c {
			return false
		}
	}

	s.current += len(chars)
	return true
}

func (s *Scanner) addToken(Type TokenType, literal interface{}) {
	text := s.Source[s.start:s.current]
	s.Tokens = append(s.Tokens, Token{
		Type:    Type,
		Lexeme:  text,
		Literal: literal,
		Line:    s.line,
	})
}
