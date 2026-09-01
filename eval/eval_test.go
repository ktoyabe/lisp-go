package eval

import (
	"fmt"
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

func testEvalObject(t *testing.T, input string, env *env.Env, want object.Object) string {
	got, err := eval(t, input, env)
	if err != nil {
		return fmt.Sprintf("eval has error. error=%v", err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		return fmt.Sprintf("input: %s, want=%v, got=%v, diff=%v", input, want, got, diff)
	}
	return ""
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
		if diff := testEvalObject(t, tt, env, wants[i]); diff != "" {
			t.Error(diff)
		}
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
		if diff := testEvalObject(t, tt, env, wants[i]); diff != "" {
			t.Error(diff)
		}
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
		if diff := testEvalObject(t, tt, env, wants[i]); diff != "" {
			t.Error(diff)
		}
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
		if diff := testEvalObject(t, tt, env, wants[i]); diff != "" {
			t.Error(diff)
		}
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
		if diff := testEvalObject(t, tt, env, wants[i]); diff != "" {
			t.Error(diff)
		}
	}
}

func TestEvalOr(t *testing.T) {
	env := env.New()

	inputs := []string{
		"(| #t #t)",
		"(| #t #f)",
		"(| #f #f)",
	}
	wants := []object.Object{
		&object.BoolObject{Value: true},
		&object.BoolObject{Value: true},
		&object.BoolObject{Value: false},
	}

	for i, tt := range inputs {
		if diff := testEvalObject(t, tt, env, wants[i]); diff != "" {
			t.Error(diff)
		}
	}
}

func TestEvalAnd(t *testing.T) {
	env := env.New()

	inputs := []string{
		"(& #t #t)",
		"(& #t #f)",
		"(& #f #f)",
	}
	wants := []object.Object{
		&object.BoolObject{Value: true},
		&object.BoolObject{Value: false},
		&object.BoolObject{Value: false},
	}

	for i, tt := range inputs {
		if diff := testEvalObject(t, tt, env, wants[i]); diff != "" {
			t.Error(diff)
		}
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
		if diff := testEvalObject(t, tt, env, wants[i]); diff != "" {
			t.Error(diff)
		}
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
		if diff := testEvalObject(t, input, env, wants[i]); diff != "" {
			t.Error(diff)
		}
	}
}

func TestEvalFib(t *testing.T) {
	env := env.New()

	inputs := []string{
		"(define fib (lambda (n) (if (< n 2) 1 (+ (fib (- n 1)) (fib (- n 2))))))",
		"(fib 1)",
		"(fib 2)",
		"(fib 3)",
		"(fib 4)",
		"(fib 5)",
	}

	wants := []object.Object{
		object.VoidObject{},
		&object.IntObject{Value: 1},
		&object.IntObject{Value: 2},
		&object.IntObject{Value: 3},
		&object.IntObject{Value: 5},
		&object.IntObject{Value: 8},
	}

	for i, input := range inputs {
		if diff := testEvalObject(t, input, env, wants[i]); diff != "" {
			t.Error(diff)
		}
	}
}

func TestEvalStringEq(t *testing.T) {
	env := env.New()

	inputs := []string{
		"(= \"abc\" \"abcd\")",
		"(!= \"abc\" \"abcd\")",
	}

	wants := []object.Object{
		&object.BoolObject{Value: false},
		&object.BoolObject{Value: true},
	}

	for i, input := range inputs {
		if diff := testEvalObject(t, input, env, wants[i]); diff != "" {
			t.Error(diff)
		}
	}
}
