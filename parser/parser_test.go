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

func TestParseRecursive(t *testing.T) {
	input := `
	(+ 5 (* 2 3))
	`
	l := lexer.New(input)
	p := New(l)

	got, err := p.Parse()
	if err != nil {
		t.Errorf("parse failed. error=%s", err)
	}

	want := &object.ListObject{
		Value: []object.Object{
			&object.SymbolObject{Value: "+"},
			&object.IntObject{Value: 5},
			&object.ListObject{Value: []object.Object{
				&object.SymbolObject{Value: "*"},
				&object.IntObject{Value: 2},
				&object.IntObject{Value: 3},
			}},
		},
	}

	for i, tt := range want.Value {
		testObject(t, tt, got.Value[i])
	}
}

func testObject(t *testing.T, want object.Object, got object.Object) {
	want_list, ok := want.(*object.ListObject)
	if ok { // expected is list
		got_list, ok := got.(*object.ListObject)
		if !ok {
			t.Errorf("got not *object.ListObject. got=%T", got)
		}
		for i, want_obj := range want_list.Value {
			if diff := cmp.Diff(got_list.Value[i], want_obj); diff != "" {
				t.Errorf("[%v] want: %v, got=%v", i, want_obj, got_list.Value[i])
			}
		}
	} else { // expected is not list
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("want: %v, got=%v", want, got)
		}
	}

}
