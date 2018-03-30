package parser

import (
	"testing"
)

func TestAdvance(t *testing.T) {
	s := "100"
	tokens, err := Scan(s)
	if err != nil {
		t.Fatalf("Error scanning: %v\n", err)
	}
	p := &Parser{tokens: tokens, lookahead: nil}

	p.advance()
	if p.lookahead.Value != "100" {
		t.Errorf("advance, expected: %s, got: %s", "100", p.lookahead.Value)
	}

	p.advance()
	if p.lookahead.Kind != End {
		t.Errorf("advance: expected: %v, got: %v", End, p.lookahead.Kind)
	}
}

func TestParserSimple(t *testing.T) {
	s := "100 + 200 * (1 + 2 + 3 * (1 + 2))"
	tokens, err := Scan(s)
	if err != nil {
		t.Fatalf("Error scanning: %v\n", err)
	}
	// printTokens(tokens)
	p := &Parser{tokens: tokens, lookahead: nil}

	p.Parse()
}

func TestParserFail(t *testing.T) {
	s := "(100"
	tokens, err := Scan(s)
	if err != nil {
		t.Fatalf("Error scanning: %v\n", err)
	}
	// printTokens(tokens)
	p := &Parser{tokens: tokens, lookahead: nil}

	myerr := p.Parse()
	if myerr == nil {
		t.Errorf("Expected error but was %v", myerr)
	}
}
