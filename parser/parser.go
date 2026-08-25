package parser

import (
	"fmt"
	"lisp-go/lexer"
	"lisp-go/object"
	"lisp-go/token"
	"strconv"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.NextToken()
	p.NextToken()

	return p
}

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Parse() (*object.ListObject, error) {
	tok := p.curToken
	if tok.Type != token.LPAREN {
		return nil, fmt.Errorf("expected LParent. found %v", tok)
	}

	p.NextToken()
	return p.parseList()
}

func (p *Parser) parseList() (*object.ListObject, error) {
	var list []object.Object

	for {
		tok := p.curToken
		if tok.Type == token.EOF {
			return nil, fmt.Errorf("insufficient token.")
		}

		switch tok.Type {
		case token.PLUS:
			o := &object.SymbolObject{Value: tok.Literal}
			list = append(list, o)
		case token.INT:
			i, err := strconv.Atoi(tok.Literal)
			if err != nil {
				return nil, fmt.Errorf("failed to Atoi. value=%v", tok.Literal)
			}
			o := &object.IntObject{Value: i}
			list = append(list, o)
		case token.RPAREN:
			p.NextToken()
			return &object.ListObject{Value: list}, nil
		default:
			return nil, fmt.Errorf("Unsupported token. type=%v, literal=%v", tok.Type, tok.Literal)
		}
		p.NextToken()
	}

}
