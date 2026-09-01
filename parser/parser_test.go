package parser

import (
	"fmt"
	"lisp-go/lexer"
	"lisp-go/object"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	input := "(+ 5 10)"
	want := []object.Object{
		&object.OperatorObject{Value: "+"},
		&object.IntObject{Value: 5},
		&object.IntObject{Value: 10},
	}
	if diff := testParse(t, input, want); diff != "" {
		t.Error(diff)
	}
}

func TestParseMinus(t *testing.T) {
	input := "(- -5 10)"
	want := []object.Object{
		&object.OperatorObject{Value: "-"},
		&object.IntObject{Value: -5},
		&object.IntObject{Value: 10},
	}
	if diff := testParse(t, input, want); diff != "" {
		t.Error(diff)
	}
}

func TestParseLess(t *testing.T) {
	input := "(< 5 10)"
	want := []object.Object{
		&object.OperatorObject{Value: "<"},
		&object.IntObject{Value: 5},
		&object.IntObject{Value: 10},
	}
	if diff := testParse(t, input, want); diff != "" {
		t.Error(diff)
	}
}

func TestParseEQ(t *testing.T) {
	input := "(= 5 10)"
	want := []object.Object{
		&object.OperatorObject{Value: "="},
		&object.IntObject{Value: 5},
		&object.IntObject{Value: 10},
	}
	if diff := testParse(t, input, want); diff != "" {
		t.Error(diff)
	}
}

func TestParseNotEQ(t *testing.T) {
	input := "(!= 5 10)"
	want := []object.Object{
		&object.OperatorObject{Value: "!="},
		&object.IntObject{Value: 5},
		&object.IntObject{Value: 10},
	}
	if diff := testParse(t, input, want); diff != "" {
		t.Error(diff)
	}
}

func TestParseOr(t *testing.T) {
	input := "(| #t #f)"
	want := []object.Object{
		&object.OperatorObject{Value: "|"},
		&object.BoolObject{Value: true},
		&object.BoolObject{Value: false},
	}
	if diff := testParse(t, input, want); diff != "" {
		t.Error(diff)
	}
}

func TestParseAnd(t *testing.T) {
	input := "(& #t #f)"
	want := []object.Object{
		&object.OperatorObject{Value: "&"},
		&object.BoolObject{Value: true},
		&object.BoolObject{Value: false},
	}
	if diff := testParse(t, input, want); diff != "" {
		t.Error(diff)
	}
}

func TestParseRecursive(t *testing.T) {
	input := `
	(+ 5 (* 2 3))
	`
	want := []object.Object{
		&object.OperatorObject{Value: "+"},
		&object.IntObject{Value: 5},
		&object.ListObject{Value: []object.Object{
			&object.OperatorObject{Value: "*"},
			&object.IntObject{Value: 2},
			&object.IntObject{Value: 3},
		}},
	}
	if diff := testParse(t, input, want); diff != "" {
		t.Error(diff)
	}
}

func testObject(t *testing.T, want object.Object, got object.Object) string {
	switch w := want.(type) {
	case *object.ListObject:
		return testListObject(t, w, got)
	case *object.LambdaObject:
		return testLambdaObject(t, w, got)
	default:
		if diff := cmp.Diff(want, got); diff != "" {
			return diff
		}
	}
	return ""
}

func TestParseSymbol(t *testing.T) {
	input := "(define a 10)"
	want := []object.Object{
		&object.SymbolObject{Value: "define"},
		&object.SymbolObject{Value: "a"},
		&object.IntObject{Value: 10},
	}
	if diff := testParse(t, input, want); diff != "" {
		t.Error(diff)
	}
}

func TestParseIf(t *testing.T) {
	input := "(if #t 3 1)"
	want := []object.Object{
		&object.SymbolObject{Value: "if"},
		&object.BoolObject{Value: true},
		&object.IntObject{Value: 3},
		&object.IntObject{Value: 1},
	}
	if diff := testParse(t, input, want); diff != "" {
		t.Error(diff)
	}
}

func testListObject(_ *testing.T, want *object.ListObject, got object.Object) string {
	got_list, ok := got.(*object.ListObject)
	if !ok {
		return fmt.Sprintf("got not *object.ListObject. got=%T", got)
	}
	for i, o := range want.Value {
		if diff := cmp.Diff(got_list.Value[i], o); diff != "" {
			return fmt.Sprintf("[%v] want: %v, got=%v", i, o, got_list.Value[i])
		}
	}
	return ""
}

func testLambdaObject(_ *testing.T, want *object.LambdaObject, got object.Object) string {
	lambda, ok := got.(*object.LambdaObject)
	if !ok {
		return fmt.Sprintf("got not *object.LambdaObject. got=%T", got)
	}

	if len(want.Params) != len(lambda.Params) {
		return fmt.Sprintf("LambdaObject#Params len different. want=%d, got=%d", len(want.Params), len(lambda.Params))
	}
	for i, _ := range want.Params {
		if diff := cmp.Diff(lambda.Params[i], want.Params[i]); diff != "" {
			return fmt.Sprintf("Params[%d] want=%v, got=%v", i, want.Params[i], lambda.Params[i])
		}
	}

	if len(want.Body.Value) != len(lambda.Body.Value) {
		return fmt.Sprintf("LambdaObject#Body len different. want=%d, got=%d", len(want.Body.Value), len(lambda.Body.Value))
	}
	for i, _ := range want.Body.Value {
		if diff := cmp.Diff(lambda.Body.Value[i], want.Body.Value[i]); diff != "" {
			return fmt.Sprintf("[%d] want: %v, got=%v", i, want.Body.Value[i], lambda.Body.Value[i])
		}
	}
	return ""
}

func testParse(t *testing.T, input string, want []object.Object) string {
	l := lexer.New(input)
	p := New(l)

	got, err := p.Parse()
	if err != nil {
		return fmt.Sprintf("parse failed. error=%s", err)
	}

	for i, want := range want {
		if diff := testObject(t, want, got.Value[i]); diff != "" {
			return diff
		}
	}
	return ""
}

func TestParseLambda(t *testing.T) {
	input := "(define add (lambda (x y) (+ x y)))"
	want := []object.Object{
		&object.SymbolObject{Value: "define"},
		&object.SymbolObject{Value: "add"},
		&object.LambdaObject{
			Params: []string{"x", "y"},
			Body: &object.ListObject{Value: []object.Object{
				&object.OperatorObject{Value: "+"},
				&object.SymbolObject{Value: "x"},
				&object.SymbolObject{Value: "y"},
			}},
		},
	}

	testParse(t, input, want)
}
