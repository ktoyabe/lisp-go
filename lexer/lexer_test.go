package lexer

import (
	"lisp-go/token"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func testLexer(t *testing.T, input string, want []token.Token) string {
	l := New(input)

	for _, tt := range want {
		tok := l.NextToken()

		if diff := cmp.Diff(tok, tt); diff != "" {
			return diff
		}
	}
	return ""
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

	if diff := testLexer(t, input, tests); diff != "" {
		t.Error(diff)
	}
}

func TestNextTokenMinus(t *testing.T) {
	input := `
	(- -5 10)
	`
	tests := []token.Token{
		{Type: token.LPAREN, Literal: "("},
		{Type: token.MINUS, Literal: "-"},
		{Type: token.INT, Literal: "-5"},
		{Type: token.INT, Literal: "10"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.EOF, Literal: ""},
	}

	if diff := testLexer(t, input, tests); diff != "" {
		t.Error(diff)
	}
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

	if diff := testLexer(t, input, tests); diff != "" {
		t.Error(diff)
	}
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

	if diff := testLexer(t, input, tests); diff != "" {
		t.Error(diff)
	}
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

	if diff := testLexer(t, input, tests); diff != "" {
		t.Error(diff)
	}
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

	if diff := testLexer(t, input, tests); diff != "" {
		t.Error(diff)
	}
}

func TestNextTokenNotEq(t *testing.T) {
	input := "(!= 3 1)"

	tests := []token.Token{
		{Type: token.LPAREN, Literal: "("},
		{Type: token.NOT_EQ, Literal: "!="},
		{Type: token.INT, Literal: "3"},
		{Type: token.INT, Literal: "1"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.EOF, Literal: ""},
	}

	if diff := testLexer(t, input, tests); diff != "" {
		t.Error(diff)
	}
}

func TestNextTokenLambda(t *testing.T) {
	input := "(lambda (x y) (+ x y))"

	tests := []token.Token{
		{Type: token.LPAREN, Literal: "("},
		{Type: token.LAMBDA, Literal: "lambda"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.SYMBOL, Literal: "x"},
		{Type: token.SYMBOL, Literal: "y"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.SYMBOL, Literal: "x"},
		{Type: token.SYMBOL, Literal: "y"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.EOF, Literal: ""},
	}

	if diff := testLexer(t, input, tests); diff != "" {
		t.Error(diff)
	}
}

func TestBooleanAnd(t *testing.T) {
	input := "(& #t #f)"

	tests := []token.Token{
		{Type: token.LPAREN, Literal: "("},
		{Type: token.AND, Literal: "&"},
		{Type: token.TRUE, Literal: "#t"},
		{Type: token.FALSE, Literal: "#f"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.EOF, Literal: ""},
	}

	if diff := testLexer(t, input, tests); diff != "" {
		t.Error(diff)
	}
}

func TestBooleanOr(t *testing.T) {
	input := "(| #t #f)"

	tests := []token.Token{
		{Type: token.LPAREN, Literal: "("},
		{Type: token.OR, Literal: "|"},
		{Type: token.TRUE, Literal: "#t"},
		{Type: token.FALSE, Literal: "#f"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.EOF, Literal: ""},
	}

	if diff := testLexer(t, input, tests); diff != "" {
		t.Error(diff)
	}
}

func TestStringEq(t *testing.T) {
	input := `(= "abc" "efgh")`

	tests := []token.Token{
		{Type: token.LPAREN, Literal: "("},
		{Type: token.EQ, Literal: "="},
		{Type: token.STRING, Literal: "abc"},
		{Type: token.STRING, Literal: "efgh"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.EOF, Literal: ""},
	}

	if diff := testLexer(t, input, tests); diff != "" {
		t.Error(diff)
	}
}
