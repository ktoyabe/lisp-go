package lexer

import (
	"lisp-go/token"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNextToken(t *testing.T) {
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

	l := New(input)

	for _, tt := range tests {
		tok := l.NextToken()

		if diff := cmp.Diff(tok, tt); diff != "" {
			t.Errorf("actual: %v, expected: %v", tok, tt)
		}
	}

}
