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
func (p *Parser) Parse() (err interface{}) {
	defer func() {
		if err = recover(); err != nil {
			return
		}
	}()
	p.advance()
	p.exp()
	return
}

func (p *Parser) exp() {
	p.term()
	p.expD()
}

func (p *Parser) expD() {
	fmt.Printf("in expD, lookahead: %v\n", p.lookahead)
	if p.lookahead.Value == "+" {
		p.match(&Token{Kind: Operator, Value: "+"})
		p.exp()
		return
	}
	if p.lookahead.Value == "-" {
		p.match(&Token{Kind: Operator, Value: "-"})
		p.exp()
		return
	}
	if p.lookahead.Kind == End {
		return
	}
}

func (p *Parser) term() {
	p.factor()
	p.termD()
}

func (p *Parser) termD() {
	fmt.Printf("in termD, lookahead: %v\n", p.lookahead)
	if p.lookahead.Value == "*" {
		p.match(&Token{Kind: Operator, Value: "*"})
		p.term()
		return
	}
	if p.lookahead.Value == "/" {
		p.match(&Token{Kind: Operator, Value: "/"})
		p.term()
		return
	}
	if p.lookahead.Kind == End {
		return
	}
}

func (p *Parser) factor() {
	fmt.Printf("in factor, lookahead: %v\n", p.lookahead)
	if p.lookahead.Kind == Num {
		p.match(p.lookahead)
		return
	}
	if p.lookahead.Value == "(" {
		p.match(&Token{Kind: Paren, Value: "("})
		p.exp()
		p.match(&Token{Kind: Paren, Value: ")"})
		return
	}
}

func (p *Parser) advance() {
	if len(p.tokens) < 1 {
		p.lookahead = &Token{Kind: End, Value: ""}
		p.tokens = make([]*Token, 0)
		return
	}
	p.lookahead, p.tokens = p.tokens[0], p.tokens[1:]
}

func (p *Parser) match(tok *Token) {
	if tok.Kind == Num {
		p.advance()
		return
	}
	if tok.Value != p.lookahead.Value {
		panic(fmt.Sprintf("error lookahead match; lookahead: %s, current: %v", p.lookahead.Value, tok))
	}
	p.advance()
}
