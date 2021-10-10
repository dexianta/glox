package errorhandle

import (
	"dexianta/glox/scanner"
	"fmt"
)

var HadError = false

func report(line int, where, msg string) {
	fmt.Printf("[line \"%d\"] Error %s \": \" %s\n", line, where, msg)
	HadError = true
}

func LoxError(token scanner.Token, msg string) {
	if token.Type == scanner.EOF {
		report(token.Line, " at end", msg)
	} else {
		report(token.Line, "at '"+token.Lexeme+"'", msg)
	}
}

func ScanError(line int, msg string) {
	report(line, "", msg)
}
