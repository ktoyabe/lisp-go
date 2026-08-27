package parser

import (
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
	testParse(t, input, want)
}

func TestParseLess(t *testing.T) {
	input := "(< 5 10)"
	want := []object.Object{
		&object.OperatorObject{Value: "<"},
		&object.IntObject{Value: 5},
		&object.IntObject{Value: 10},
	}
	testParse(t, input, want)
}

func TestParseEQ(t *testing.T) {
	input := "(= 5 10)"
	want := []object.Object{
		&object.OperatorObject{Value: "="},
		&object.IntObject{Value: 5},
		&object.IntObject{Value: 10},
	}
	testParse(t, input, want)
}

func TestParseNotEQ(t *testing.T) {
	input := "(!= 5 10)"
	want := []object.Object{
		&object.OperatorObject{Value: "!="},
		&object.IntObject{Value: 5},
		&object.IntObject{Value: 10},
	}
	testParse(t, input, want)
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
	testParse(t, input, want)
}

func testObject(t *testing.T, want object.Object, got object.Object) {
	switch w := want.(type) {
	case *object.ListObject:
		testListObject(t, w, got)
	case *object.LambdaObject:
		testLambdaObject(t, w, got)
	default:
		if diff := cmp.Diff(want, got); diff != "" {
			t.Error(diff)
		}
	}
}

func TestParseSymbol(t *testing.T) {
	input := "(define a 10)"
	want := []object.Object{
		&object.SymbolObject{Value: "define"},
		&object.SymbolObject{Value: "a"},
		&object.IntObject{Value: 10},
	}
	testParse(t, input, want)
}

func TestParseIf(t *testing.T) {
	input := "(if #t 3 1)"
	want := []object.Object{
		&object.SymbolObject{Value: "if"},
		&object.BoolObject{Value: true},
		&object.IntObject{Value: 3},
		&object.IntObject{Value: 1},
	}
	testParse(t, input, want)
}

func testListObject(t *testing.T, want *object.ListObject, got object.Object) {
	got_list, ok := got.(*object.ListObject)
	if !ok {
		t.Errorf("got not *object.ListObject. got=%T", got)
	}
	for i, o := range want.Value {
		if diff := cmp.Diff(got_list.Value[i], o); diff != "" {
			t.Errorf("[%v] want: %v, got=%v", i, o, got_list.Value[i])
		}
	}

}

func testLambdaObject(t *testing.T, want *object.LambdaObject, got object.Object) {
	lambda, ok := got.(*object.LambdaObject)
	if !ok {
		t.Errorf("got not *object.LambdaObject. got=%T", got)
	}

	if len(want.Params) != len(lambda.Params) {
		t.Errorf("LambdaObject#Params len different. want=%d, got=%d", len(want.Params), len(lambda.Params))
	}
	for i, _ := range want.Params {
		if diff := cmp.Diff(lambda.Params[i], want.Params[i]); diff != "" {
			t.Errorf("Params[%d] want=%v, got=%v", i, want.Params[i], lambda.Params[i])
		}
	}

	if len(want.Body.Value) != len(lambda.Body.Value) {
		t.Errorf("LambdaObject#Body len different. want=%d, got=%d", len(want.Body.Value), len(lambda.Body.Value))
	}
	for i, _ := range want.Body.Value {
		if diff := cmp.Diff(lambda.Body.Value[i], want.Body.Value[i]); diff != "" {
			t.Errorf("[%d] want: %v, got=%v", i, want.Body.Value[i], lambda.Body.Value[i])
		}
	}
}

func testParse(t *testing.T, input string, want []object.Object) {
	l := lexer.New(input)
	p := New(l)

	got, err := p.Parse()
	if err != nil {
		t.Errorf("parse failed. error=%s", err)
	}

	for i, tt := range want {
		testObject(t, tt, got.Value[i])
	}
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
