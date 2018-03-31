package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gerlacdt/arithmetic-parser/pkg/parser"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter expression:\n")
	expression, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("error during reading: %v", err))
	}

	result, err := eval(expression)
	if err != nil {
		panic(fmt.Sprintf("error during reading: %v", err))
	}
	fmt.Printf("result: %d\n", result)
}

func eval(expression string) (int, error) {
	tokens, err := parser.Scan(expression)
	if err != nil {
		return 0, err
	}
	p := &parser.Parser{Tokens: tokens, Lookahead: nil}
	result, err := p.Parse()
	if err != nil {
		return 0, err
	}
	return result, nil
}
