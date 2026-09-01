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

	return p.parseList()
}

// format: (x y z)
func (p *Parser) parseLambdaParams() ([]string, error) {
	tok := p.curToken
	if tok.Type != token.LPAREN {
		return nil, fmt.Errorf("parseLambdaParams: not LParent. token=%v", tok)
	}
	p.NextToken() // consume LParent

	params := []string{}
	for {
		tok := p.curToken
		if tok.Type == token.RPAREN {
			p.NextToken() // consume RParen
			return params, nil
		}
		params = append(params, tok.Literal)

		p.NextToken()
		if tok.Type == token.EOF {
			return nil, fmt.Errorf("parseLambdaParams: insufficient token. RParent is not found.")
		}
	}
}

func (p *Parser) parseLambda() (*object.LambdaObject, error) {
	// parse "lambda"
	tok := p.curToken
	if tok.Type != token.LAMBDA {
		return nil, fmt.Errorf("parseLambda: first element not lambda. token=%v", tok)
	}
	p.NextToken() // consume lambda

	// parse params
	params, err := p.parseLambdaParams()
	if err != nil {
		return nil, err
	}

	// parse body
	body, err := p.parseList()
	if err != nil {
		return nil, err
	}
	return &object.LambdaObject{Params: params, Body: body}, nil
}

func (p *Parser) parseList() (*object.ListObject, error) {
	var list []object.Object

	for {
		p.NextToken() // consume LParent

		tok := p.curToken
		if tok.Type == token.EOF {
			return nil, fmt.Errorf("insufficient token.")
		}

		switch tok.Type {
		case token.TRUE:
			o := &object.BoolObject{Value: true}
			list = append(list, o)
		case token.FALSE:
			o := &object.BoolObject{Value: false}
			list = append(list, o)
		case token.PLUS:
			o := &object.OperatorObject{Value: tok.Literal}
			list = append(list, o)
		case token.MINUS:
			o := &object.OperatorObject{Value: tok.Literal}
			list = append(list, o)
		case token.MUL:
			o := &object.OperatorObject{Value: tok.Literal}
			list = append(list, o)
		case token.LESS:
			o := &object.OperatorObject{Value: tok.Literal}
			list = append(list, o)
		case token.GREATER:
			o := &object.OperatorObject{Value: tok.Literal}
			list = append(list, o)
		case token.EQ:
			o := &object.OperatorObject{Value: tok.Literal}
			list = append(list, o)
		case token.NOT_EQ:
			o := &object.OperatorObject{Value: tok.Literal}
			list = append(list, o)
		case token.AND:
			o := &object.OperatorObject{Value: tok.Literal}
			list = append(list, o)
		case token.OR:
			o := &object.OperatorObject{Value: tok.Literal}
			list = append(list, o)
		case token.INT:
			i, err := strconv.Atoi(tok.Literal)
			if err != nil {
				return nil, fmt.Errorf("failed to Atoi. value=%v", tok.Literal)
			}
			o := &object.IntObject{Value: i}
			list = append(list, o)
		case token.DEFINE:
			o := &object.SymbolObject{Value: tok.Literal}
			list = append(list, o)
		case token.LPAREN:
			if p.peekToken.Type == token.LAMBDA {
				p.NextToken() // consume '('
				lambda, err := p.parseLambda()
				if err != nil {
					return nil, err
				}
				list = append(list, lambda)

			} else {
				subList, err := p.parseList()
				if err != nil {
					return nil, err
				}
				list = append(list, subList)
			}
		case token.RPAREN:
			return &object.ListObject{Value: list}, nil
		default:
			o := &object.SymbolObject{Value: tok.Literal}
			list = append(list, o)
			// return nil, fmt.Errorf("Unsupported token. type=%v, literal=%v", tok.Type, tok.Literal)
		}
	}

}
