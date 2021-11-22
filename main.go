package main

import (
	"bufio"
	"dexianta/glox/errorhandle"
	"dexianta/glox/parser"
	"dexianta/glox/scanner"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		panic("Usage: glox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}

func runFile(path string) error {
	contentBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return run(string(contentBytes))
}

func runPrompt() error {
	for {
		fmt.Printf("> ")
		reader := bufio.NewReader(os.Stdin)
		//TODO: multi-line input
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("errorhandle reading line: ", err.Error())
		}

		if err := run(line); err != nil {
			fmt.Println("errorhandle running line: ", err.Error())
		}
	}
}

func run(code string) error {
	s := scanner.NewScanner(code)
	tokens := s.ScanTokens()
	parser := parser.NewParser(tokens)
	_ = parser.Parse()

	if errorhandle.HadError {
		return errors.New("something is wrong")
	}

	return nil
}
