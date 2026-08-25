package parser

import (
	"lisp-go/lexer"
	"lisp-go/object"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	input := `
	(+ 5 10)
	`
	l := lexer.New(input)
	p := New(l)

	actual, err := p.Parse()
	if err != nil {
		t.Errorf("parse failed. error=%s", err)
	}

	expected := []object.Object{
		&object.SymbolObject{Value: "+"},
		&object.IntObject{Value: 5},
		&object.IntObject{Value: 10},
	}

	for i, tt := range expected {
		want := actual.Value[i]
		if diff := cmp.Diff(want, tt); diff != "" {
			t.Errorf("%v, actual: %v, expected: %v", i, want, tt)
		}
	}
}
