package parser

import (
	"fmt"
)

// <Exp> ::= <Term> <Exp'>
// <Exp'> ::= + <Exp> | - <Exp> | ε

// <Term> ::= <Factor> <Term'>
// <Term'> ::= * <Term> | / <Term> | ε

// <Factor> ::= number
// 	( <Exp> ) |
// 	- <Factor>

// Parser parses the given tokens
type Parser struct {
	tokens    []*Token
	lookahead *Token
}

// Parse the given tokens and return result of arithmetic expression
func (p *Parser) Parse() {
	p.advance()
	p.exp()
}

func (p *Parser) exp() {
	p.term()
	p.expD()
}

func (p *Parser) expD() {
	if p.lookahead.Value == "+" {
		p.match("+")
		p.exp()
		return
	}
	if p.lookahead.Value == "-" {
		p.match("-")
		p.exp()
		return
	}
	return
}

func (p *Parser) term() {
	p.factor()
	p.termD()
}

func (p *Parser) termD() {
	if p.lookahead.Value == "*" {
		p.match("*")
		p.term()
		return
	}
	if p.lookahead.Value == "/" {
		p.match("/")
		p.term()
		return
	}
	return
}

func (p *Parser) factor() {
	if p.lookahead.Kind == Num {
		// value, err := strconv.Atoi(p.lookahead.Value)
		// if err != nil {
		// 	panic(fmt.Sprintf("Error converting number token: %s", p.lookahead.Value))
		// }
		return
	}
	if p.lookahead.Value == "(" {
		p.exp()
		p.match(")")
		return
	}

}

func (p *Parser) advance() {
	if len(p.tokens) < 1 {
		p.lookahead = nil
		p.tokens = make([]*Token, 0)
		return
	}
	p.lookahead, p.tokens = p.tokens[0], p.tokens[1:]
}

func (p *Parser) match(tok string) {
	if tok != p.lookahead.Value {
		panic(fmt.Sprintf("error lookahead match; lookahead: %s, current: %s", p.lookahead.Value, tok))
	}
	p.advance()
}
