package lexer

import (
	"lisp-go/token"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func testLexer(t *testing.T, input string, want []token.Token) {
	l := New(input)

	for _, tt := range want {
		tok := l.NextToken()

		if diff := cmp.Diff(tok, tt); diff != "" {
			t.Errorf("got: %v, want: %v", tok, tt)
		}
	}
}

func TestNextTokenINT(t *testing.T) {
	input := `
	(+ 5 10)
	`
	tests := []token.Token{
		{Type: token.LPAREN, Literal: "("},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.INT, Literal: "5"},
		{Type: token.INT, Literal: "10"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.EOF, Literal: ""},
	}

	testLexer(t, input, tests)
}

func TestNextTokenSYMBOL(t *testing.T) {
	input := `
	(define mul (* x y))
	`
	tests := []token.Token{
		{Type: token.LPAREN, Literal: "("},
		{Type: token.DEFINE, Literal: "define"},
		{Type: token.SYMBOL, Literal: "mul"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.MUL, Literal: "*"},
		{Type: token.SYMBOL, Literal: "x"},
		{Type: token.SYMBOL, Literal: "y"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.EOF, Literal: ""},
	}

	testLexer(t, input, tests)
}

func TestNextTokenIf(t *testing.T) {
	input := "(if #t #f #t)"

	tests := []token.Token{
		{Type: token.LPAREN, Literal: "("},
		{Type: token.IF, Literal: "if"},
		{Type: token.TRUE, Literal: "#t"},
		{Type: token.FALSE, Literal: "#f"},
		{Type: token.TRUE, Literal: "#t"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.EOF, Literal: ""},
	}

	testLexer(t, input, tests)
}

func TestNextTokenLess(t *testing.T) {
	input := "(< 3 1)"

	tests := []token.Token{
		{Type: token.LPAREN, Literal: "("},
		{Type: token.LESS, Literal: "<"},
		{Type: token.INT, Literal: "3"},
		{Type: token.INT, Literal: "1"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.EOF, Literal: ""},
	}

	testLexer(t, input, tests)
}

func TestNextTokenEq(t *testing.T) {
	input := "(= 3 1)"

	tests := []token.Token{
		{Type: token.LPAREN, Literal: "("},
		{Type: token.EQ, Literal: "="},
		{Type: token.INT, Literal: "3"},
		{Type: token.INT, Literal: "1"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.EOF, Literal: ""},
	}

	testLexer(t, input, tests)
}
