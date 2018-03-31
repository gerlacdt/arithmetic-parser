package parser

import (
	"errors"
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

func TestParserTable(t *testing.T) {
	// Table parser test
	tt := []struct {
		name   string
		input  string
		result int
		err    error
	}{
		{"1", "1", 1, nil},
		{"1 + 2", "1 + 2", 3, nil},
		{"1 * 2", "1 * 2", 2, nil},
		{"(1 + 2)", "(1 + 2)", 3, nil},
		{"(1 * 2)", "(1 * 2)", 2, nil},
		{"(1 + 2) + 1", "(1 + 2) + 1", 4, nil},
		{"3 + 4 * 4", "3 + 4 * 4", 19, nil},
		{"(1 + 2) * (3 * 3) + 2", "(1 + 2) * (3 * 3) + 2", 29, nil},
		{"(1 + 2 * (3 + 4) + 1) + 1", "(1 + 2 * (3 + 4) + 1) + 1", 17, nil},
		{"1 + 2 with whitespaces", "  1      +      2    ", 3, nil},
		{"(100", "(100", 0, errors.New("error lookahead match; lookahead: , current: &{2 )}")},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tokens, err := Scan(tc.input)
			if err != nil {
				t.Fatalf("Error scanning: %v\n", err)
			}
			// printTokens(tokens)
			p := &Parser{tokens: tokens, lookahead: nil}
			result, err := p.Parse()
			if err != nil {
				if err.Error() != tc.err.Error() {
					t.Errorf("eval of %s, expected error: %v, got: %v", tc.name, tc.err, err)
				}
			}
			if result != tc.result {
				t.Errorf("eval of %s, expected: %d, got: %d", tc.name, tc.result, result)
			}
		})
	}
}
