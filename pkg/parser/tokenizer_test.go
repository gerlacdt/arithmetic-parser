package parser

import (
	"fmt"
	"testing"
)

func printTokens(tokens []*Token) {
	for _, tok := range tokens {
		fmt.Printf("tok: %v\n", tok)
	}
}

func TestScanSuccess(t *testing.T) {
	s := "100 + ( 99 * 2 + 1 )"
	tokens, err := Scan(s)
	if err != nil {
		t.Fatalf("Error scanning: %v\n", err)
	}
	// printTokens(tokens)
	expectedLength := 9
	if len(tokens) != expectedLength {
		t.Errorf("tokens length, expected: %d, got: %d", expectedLength, len(tokens))
	}
}

func TestScanSyntaxError(t *testing.T) {
	s := "100 + a + ( 99 * 2 + 1 )"
	_, err := Scan(s)
	if err == nil {
		t.Fatal("Expected Error during scanning.\n")
	}
}
