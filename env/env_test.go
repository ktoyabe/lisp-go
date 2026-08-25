package env

import (
	"lisp-go/object"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEnvGetSuccess(t *testing.T) {
	e := New()
	e.Set("a", &object.IntObject{Value: 3})

	o, ok := e.Get("a")
	if !ok {
		t.Errorf("want=%v, got=%v", true, ok)
	}
	want := &object.IntObject{Value: 3}
	if diff := cmp.Diff(o, want); diff != "" {
		t.Error(diff)
	}
}

func TestEnvGetFail(t *testing.T) {
	e := New()
	e.Set("a", &object.IntObject{Value: 3})

	_, ok := e.Get("b")
	if ok {
		t.Errorf("want=%v, got=%v", false, ok)
	}
}

func intObj(val int) *object.IntObject {
	return &object.IntObject{Value: val}
}

func TestExtendEnvGet(t *testing.T) {
	p := New()
	p.Set("a", intObj(3))
	p.Set("b", intObj(4))

	e := Extend(p)
	e.Set("a", intObj(5))

	// overwrite
	o1, _ := e.Get("a")
	if diff := cmp.Diff(o1, intObj(5)); diff != "" {
		t.Error(diff)
	}

	// env missing, parent hit
	o2, _ := e.Get("b")
	if diff := cmp.Diff(o2, intObj(4)); diff != "" {
		t.Error(diff)
	}

	// env and parent missing
	_, ok := e.Get("c")
	if ok {
		t.Errorf("want=%v, got=%v", false, ok)
	}

}
