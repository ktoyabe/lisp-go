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

func testEvalObject(t *testing.T, input string, env *env.Env, want object.Object) {
	got, err := eval(t, input, env)
	if err != nil {
		t.Errorf("eval has error. error=%v", err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("input: %s, diff=%v", input, diff)
	}
}

func TestEvalBinaryOp(t *testing.T) {
	env := env.New()

	inputs := []string{
		"(+ 5 10)",
		"(* 2 3)",
		"(+ (+ 2 3) (+ 1 2))",
		"(- -5 10)",
	}
	wants := []object.Object{
		&object.IntObject{Value: 15},
		&object.IntObject{Value: 6},
		&object.IntObject{Value: 8},
		&object.IntObject{Value: -15},
	}

	for i, tt := range inputs {
		testEvalObject(t, tt, env, wants[i])
	}
}

func TestEval(t *testing.T) {
	env := env.New()

	inputs := []string{
		"(define abc 3)",
		"(* 2 abc)",
	}
	wants := []object.Object{
		object.VoidObject{},
		&object.IntObject{Value: 6},
		&object.BoolObject{Value: true},
	}

	for i, tt := range inputs {
		testEvalObject(t, tt, env, wants[i])
	}
}

func TestEvalComparison(t *testing.T) {
	env := env.New()

	inputs := []string{
		"(< 2 1)",
		"(< 2 2)",
		"(< 2 3)",
	}
	wants := []object.Object{
		&object.BoolObject{Value: false},
		&object.BoolObject{Value: false},
		&object.BoolObject{Value: true},
	}

	for i, tt := range inputs {
		testEvalObject(t, tt, env, wants[i])
	}
}

func TestEvalEq(t *testing.T) {
	env := env.New()

	inputs := []string{
		"(= 2 1)",
		"(= 2 2)",
		"(= 2 3)",
	}
	wants := []object.Object{
		&object.BoolObject{Value: false},
		&object.BoolObject{Value: true},
		&object.BoolObject{Value: false},
	}

	for i, tt := range inputs {
		testEvalObject(t, tt, env, wants[i])
	}
}

func TestEvalNotEq(t *testing.T) {
	env := env.New()

	inputs := []string{
		"(!= 2 1)",
		"(!= 2 2)",
		"(!= 2 3)",
	}
	wants := []object.Object{
		&object.BoolObject{Value: true},
		&object.BoolObject{Value: false},
		&object.BoolObject{Value: true},
	}

	for i, tt := range inputs {
		testEvalObject(t, tt, env, wants[i])
	}
}

func TestEvalIf(t *testing.T) {
	env := env.New()

	inputs := []string{
		"(if (< 1 2) 3 4)",
		"(if #f #f #t)",
	}
	wants := []object.Object{
		&object.IntObject{Value: 3},
		&object.BoolObject{Value: true},
	}

	for i, tt := range inputs {
		testEvalObject(t, tt, env, wants[i])
	}
}

func TestEvalLambda(t *testing.T) {
	env := env.New()

	inputs := []string{
		"(define add (lambda (x y) (+ x y)))",
		"(add 2 3)",
	}

	wants := []object.Object{
		object.VoidObject{},
		&object.IntObject{Value: 5},
	}

	for i, input := range inputs {
		testEvalObject(t, input, env, wants[i])
	}
}
