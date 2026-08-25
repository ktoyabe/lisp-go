package eval

import (
	"lisp-go/env"
	"lisp-go/lexer"
	"lisp-go/object"
	"lisp-go/parser"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func eval(t *testing.T, input string, env *env.Env) (object.Object, error) {
	l := lexer.New(input)
	p := parser.New(l)

	obj, err := p.Parse()
	if err != nil {
		t.Errorf("parse failed. error=%s", err)
	}

	return Eval(obj, env)
}

func TestEvalBinaryOp(t *testing.T) {
	env := env.New()

	inputs := []string{
		"(+ 5 10)",
	}
	wants := []object.Object{
		&object.IntObject{Value: 15},
	}

	for i, tt := range inputs {
		got, err := eval(t, tt, env)
		if err != nil {
			t.Errorf("eval has error. error=%v", err)
		}
		want := wants[i]
		if diff := cmp.Diff(want, got); diff != "" {
			t.Error(diff)
		}
	}
}
