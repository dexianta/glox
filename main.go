package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var hadError = false

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
			fmt.Println(err.Error())
		}

		if err := run(line); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func run(code string) error {
	scanner := NewScanner(code)

	tokens := scanner.scanTokens()
	for _, t := range tokens {
		fmt.Println(t)
	}

	return nil
}
