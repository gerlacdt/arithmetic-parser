package parser

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"text/scanner"
	"unicode"
)

// Token contains value and kind of token
type Token struct {
	Kind  string
	Value string
}

func getNumber(token rune, s scanner.Scanner) int {
	v := 0
	for tok := token; !unicode.IsNumber(tok); s.Scan() {
		n, err := strconv.Atoi(s.TokenText())
		if err != nil {
			log.Fatalf("Error parsing number, token: %v\n", s.TokenText())
		}
		v = v*10 + n
	}
	return v
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}

func isOperator(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	}
	return false
}

func isParen(s string) bool {
	if s == "(" || s == ")" {
		return true
	}
	return false
}

// Scan a given string and returns a slice of tokens
func Scan(src string) ([]*Token, error) {
	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	var tokens []*Token
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tstring := s.TokenText()
		if unicode.IsSpace(tok) {
			continue
		}
		if isNumber(tstring) {
			tokens = append(tokens, &Token{Kind: "num", Value: tstring})
			continue
		}
		if isOperator(tstring) {
			tokens = append(tokens, &Token{Kind: "operator", Value: tstring})
			continue
		}
		if isParen(tstring) {
			tokens = append(tokens, &Token{Kind: "paren", Value: tstring})
			continue
		}
		return nil, fmt.Errorf("Syntax error, not expected token: %s", tstring)
	}
	return tokens, nil
}
