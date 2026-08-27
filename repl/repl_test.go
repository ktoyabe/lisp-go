package repl

import (
	"bytes"
	"fmt"
	"lisp-go/object"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func toString(obj object.Object) string {
	return fmt.Sprintf("[%T] %+v", obj, obj)
}

func TestRepl(t *testing.T) {
	gots := executeRepl(t, []string{
		"(+ 1 2)",
	})
	wants := []string{
		toString(&object.IntObject{Value: 3}),
		"",
	}
	if diff := cmp.Diff(wants, gots); diff != "" {
		t.Error(diff)
	}
}

func TestReplUsingDefine(t *testing.T) {
	gots := executeRepl(t, []string{
		"(define a 3)",
		"(+ 1 a)",
	})
	wants := []string{
		toString(object.VoidObject{}),
		toString(&object.IntObject{Value: 4}),
		"",
	}
	if diff := cmp.Diff(wants, gots); diff != "" {
		t.Error(diff)
	}
}

func executeRepl(_ *testing.T, inputs []string) []string {
	input := strings.Join(inputs, "\n")
	reader := strings.NewReader(input)
	var writer bytes.Buffer

	Start(reader, &writer)

	gotStr := writer.String()
	gots := strings.Split(gotStr, "\n")

	return gots
}
