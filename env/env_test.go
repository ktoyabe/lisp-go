package env

import (
	"lisp-go/object"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEnvGetSuccess(t *testing.T) {
	e := New()
	e.set("a", &object.IntObject{Value: 3})

	o, ok := e.get("a")
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
	e.set("a", &object.IntObject{Value: 3})

	_, ok := e.get("b")
	if ok {
		t.Errorf("want=%v, got=%v", false, ok)
	}
}

func intObj(val int) *object.IntObject {
	return &object.IntObject{Value: val}
}

func TestExtendEnvGet(t *testing.T) {
	p := New()
	p.set("a", intObj(3))
	p.set("b", intObj(4))

	e := Extend(p)
	e.set("a", intObj(5))

	// overwrite
	o1, _ := e.get("a")
	if diff := cmp.Diff(o1, intObj(5)); diff != "" {
		t.Error(diff)
	}

	// env missing, parent hit
	o2, _ := e.get("b")
	if diff := cmp.Diff(o2, intObj(4)); diff != "" {
		t.Error(diff)
	}

	// env and parent missing
	_, ok := e.get("c")
	if ok {
		t.Errorf("want=%v, got=%v", false, ok)
	}

}
