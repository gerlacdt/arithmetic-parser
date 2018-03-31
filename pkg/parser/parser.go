package parser

import (
	"fmt"
	"strconv"
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
	Tokens    []*Token
	Lookahead *Token
}

// Parse the given tokens and return result of arithmetic expression
func (p *Parser) Parse() (int, error) {
	p.advance()
	return p.exp()
}

func (p *Parser) exp() (int, error) {
	val, err := p.term()
	if err != nil {
		return 0, err
	}
	result, err := p.expD(val)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (p *Parser) expD(inherited int) (int, error) {
	if p.Lookahead.Value == "+" {
		err := p.match(&Token{Kind: Operator, Value: "+"})
		if err != nil {
			return 0, err
		}
		expResult, err := p.exp()
		if err != nil {
			return 0, err
		}
		return inherited + expResult, nil
	}
	if p.Lookahead.Value == "-" {
		err := p.match(&Token{Kind: Operator, Value: "-"})
		if err != nil {
			return 0, err
		}
		expResult, err := p.exp()
		if err != nil {
			return 0, err
		}
		return inherited - expResult, nil
	}
	return inherited, nil
}

func (p *Parser) term() (int, error) {
	val, err := p.factor()
	if err != nil {
		return 0, err
	}
	result, err := p.termD(val)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (p *Parser) termD(inherited int) (int, error) {
	if p.Lookahead.Value == "*" {
		err := p.match(&Token{Kind: Operator, Value: "*"})
		if err != nil {
			return 0, err
		}
		val, err := p.term()
		if err != nil {
			return 0, err
		}
		return inherited * val, nil
	}
	if p.Lookahead.Value == "/" {
		err := p.match(&Token{Kind: Operator, Value: "/"})
		if err != nil {
			return 0, err
		}
		val, err := p.term()
		if err != nil {
			return 0, err
		}
		return inherited / val, nil
	}
	return inherited, nil
}

func (p *Parser) factor() (int, error) {
	if p.Lookahead.Kind == Num {
		val, _ := strconv.Atoi(p.Lookahead.Value)
		err := p.match(p.Lookahead)
		if err != nil {
			return 0, err
		}
		return val, nil
	}
	if p.Lookahead.Value == "(" {
		err := p.match(&Token{Kind: Paren, Value: "("})
		if err != nil {
			return 0, err
		}
		val, err := p.exp()
		if err != nil {
			return 0, err
		}
		err = p.match(&Token{Kind: Paren, Value: ")"})
		if err != nil {
			return 0, err
		}
		return val, nil
	}
	return 0, fmt.Errorf("error during parsing factor, lookahead: %v", p.Lookahead)
}

func (p *Parser) advance() {
	if len(p.Tokens) < 1 {
		p.Lookahead = &Token{Kind: End, Value: ""}
		p.Tokens = make([]*Token, 0)
		return
	}
	p.Lookahead, p.Tokens = p.Tokens[0], p.Tokens[1:]
}

func (p *Parser) match(tok *Token) error {
	if tok.Kind == Num {
		p.advance()
		return nil
	}
	if tok.Value != p.Lookahead.Value {
		return fmt.Errorf("error lookahead match; lookahead: %s, current: %v", p.Lookahead.Value, tok)
	}
	p.advance()
	return nil
}
