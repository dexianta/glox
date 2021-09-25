package main

import "fmt"

var hadError = false

func report(line int, where, msg string) {
	fmt.Printf("[line \"%d\"] Error %s \": \" %s\n", line, where, msg)
	hadError = true
}

func LoxError(token Token, msg string) {
	if token.Type == EOF {
		report(token.Line, " at end", msg)
	} else {
		report(token.Line, "at '"+token.Lexeme+"'", msg)
	}
}
