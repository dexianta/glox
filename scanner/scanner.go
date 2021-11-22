package scanner

import (
	"dexianta/glox/errorhandle"
	"fmt"
	"strconv"
)

type TokenType string

const (
	// single character tokens
	LEFT_PAREN  TokenType = "("
	RIGHT_PAREN TokenType = ")"
	LEFT_BRACE  TokenType = "{"
	RIGHT_BRACE TokenType = "}"
	COMMA       TokenType = ","
	DOT         TokenType = "."
	MINUS       TokenType = "-"
	PLUS        TokenType = "+"
	SEMICOLON   TokenType = ";"
	SLASH       TokenType = "/"
	STAR        TokenType = "*"

	// one or two character tokens
	BANG          TokenType = "!"
	BANG_EQUAL    TokenType = "!="
	EQUAL         TokenType = "="
	EQUAL_EQUAL   TokenType = "=="
	GREATER       TokenType = ">"
	GREATER_EQUAL TokenType = ">="
	LESS       TokenType = "<"
	LESS_EQUAL TokenType = "<="

	// literals
	IDENTIFIER TokenType = "identifier"
	STRING     TokenType = "string"
	NUMBER     TokenType = "number"

	// keywords
	AND    TokenType = "and"
	CLASS  TokenType = "class"
	ELSE   TokenType = "else"
	FALSE  TokenType = "false"
	FUN    TokenType = "fun"
	FOR    TokenType = "for"
	IF     TokenType = "if"
	NIL    TokenType = "nil"
	OR     TokenType = "or"
	PRINT  TokenType = "print"
	RETURN TokenType = "return"
	SUPER  TokenType = "super"
	THIS   TokenType = "this"
	TRUE   TokenType = "true"
	VAR    TokenType = "var"
	WHILE  TokenType = "while"
	EOF    TokenType = "eof"
)

var keywords = map[string]TokenType{}

func init() {
	keywords["and"] = AND
	keywords["class"] = CLASS
	keywords["else"] = ELSE
	keywords["false"] = FALSE
	keywords["for"] = FOR
	keywords["fun"] = FUN
	keywords["if"] = IF
	keywords["nil"] = NIL
	keywords["or"] = OR
	keywords["print"] = PRINT
	keywords["return"] = RETURN
	keywords["super"] = SUPER
	keywords["this"] = THIS
	keywords["true"] = TRUE
	keywords["var"] = VAR
	keywords["while"] = WHILE
}

type Token struct {
	Type    TokenType   // token type
	Lexeme  string      // the string representation
	Literal interface{} // actual value of this token
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

func (s *Scanner) ScanTokens() []Token {
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
	case '(':
		s.addToken(LEFT_PAREN, nil)
	case ')':
		s.addToken(RIGHT_PAREN, nil)
	case '{':
		s.addToken(LEFT_BRACE, nil)
	case '}':
		s.addToken(RIGHT_BRACE, nil)
	case ',':
		s.addToken(COMMA, nil)
	case '.':
		s.addToken(DOT, nil)
	case '-':
		s.addToken(MINUS, nil)
	case '+':
		s.addToken(PLUS, nil)
	case '*':
		s.addToken(STAR, nil)
	case ';':
		s.addToken(SEMICOLON, nil)

	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL, nil)
		} else {
			s.addToken(BANG, nil)
		}

	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL, nil)
		} else {
			s.addToken(EQUAL, nil)
		}

	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL, nil)
		} else {
			s.addToken(LESS, nil)
		}

	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL, nil)
		} else {
			s.addToken(GREATER, nil)
		}

	case '/':
		if s.match('/') {
			for !equalBytes(s.peek(0), []byte{'\n'}) && !s.IsAtEnd() {
				s.advance()
			}
		} else if s.match('*') {
			for !s.match('*', '/') && !s.IsAtEnd() {
				if equalBytes(s.peek(0), []byte{'\n'}) {
					s.line++
				}
				s.advance()
			}
		} else {
			s.addToken(SLASH, nil)
		}

	case ' ':
	case '\r':
	case '\t':
	case '\n':
		s.line++
	case '"':
		s.string()

	default:
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			s.identifier()
		} else {
			errorhandle.Report(s.line, "", fmt.Sprintf("Unexpected character: %c", c))
		}
	}
}

func equalBytes(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for idx := range a {
		if a[idx] != b[idx] {
			return false
		}
	}

	return true
}

func isDigit(char byte) bool {
	return char > '0' && char < '9'
}

func isAlpha(char byte) bool {
	return (char > 'a' && char < 'z') || (char > 'A' && char < 'Z') || char == '_'
}

func isAlphaNumeric(c byte) bool {
	return isDigit(c) || isAlpha(c)
}

func (s *Scanner) identifier() {
	for len(s.peek(0)) != 0 && isAlphaNumeric(s.peek(0)[0]) {
		s.advance()
	}

	text := s.Source[s.start:s.current]
	tokenType, ok := keywords[text]
	if !ok {
		tokenType = IDENTIFIER
	}

	s.addToken(tokenType, nil)
}

func (s *Scanner) number() {
	for len(s.peek(0)) != 0 && isDigit(s.peek(0)[0]) {
		s.advance()
	}

	if len(s.peek(0)) != 0 && s.peek(0)[0] == '.' && isDigit(s.peek(1)[1]) {
		s.advance() // consume the "."

		for len(s.peek(0)) != 0 && isDigit(s.peek(0)[0]) {
			s.advance()
		}
	}

	number, err := strconv.ParseFloat(s.Source[s.start:s.current], 64)
	if err != nil {
		errorhandle.Report(s.line, "",fmt.Sprintf("error handle parsing float: %s", err.Error()))
	}

	s.addToken(NUMBER, number)
}

// string by default is multiline string
func (s *Scanner) string() {
	for !equalBytes(s.peek(0), []byte{'"'}) && !s.IsAtEnd() {
		if equalBytes(s.peek(0), []byte{'\n'}) {
			s.line++
		}
		s.advance()
	}

	if s.IsAtEnd() {
		errorhandle.Report(s.line, "", "unterminated string")
		return
	}

	s.advance() // the closing "

	value := s.Source[s.start+1 : s.current-1] // remove the start & end quote
	s.addToken(STRING, value)
}

//IsAtEnd represents there's no more character left to consume
func (s *Scanner) IsAtEnd() bool {
	return s.current >= len(s.Source)
}

func (s *Scanner) IsAtEndOffset(offset int) bool {
	return s.current+offset >= len(s.Source)
}

func (s *Scanner) advance() byte {
	idx := s.current
	s.current++
	return s.Source[idx]
}

func (s *Scanner) peek(step int) (res []byte) {
	if s.IsAtEndOffset(step) {
		return
	}

	for i := s.current; i <= s.current+step; i++ {
		res = append(res, s.Source[i])
	}
	return
}

func (s *Scanner) match(chars ...byte) bool {
	if s.IsAtEnd() && s.IsAtEndOffset(len(chars)) {
		return false
	}

	for idx, c := range chars {
		if s.Source[s.current+idx] != c {
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
