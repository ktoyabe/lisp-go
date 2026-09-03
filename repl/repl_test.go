package repl

import (
	"bytes"
	"lisp-go/object"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRepl(t *testing.T) {
	gots := executeRepl(t, []string{
		"(+ 1 2)",
	})
	wants := []string{
		(&object.IntObject{Value: 3}).ToString(),
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
		(&object.IntObject{Value: 4}).ToString(),
		"",
	}
	if diff := cmp.Diff(wants, gots); diff != "" {
		t.Error(diff)
	}
}

func TestLambda(t *testing.T) {
	gots := executeRepl(t, []string{
		"(define add (lambda (x y) (+ x y)))",
	})
	wants := []string{
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
