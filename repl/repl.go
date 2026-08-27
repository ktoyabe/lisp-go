package repl

import (
	"bufio"
	"fmt"
	"io"
	"lisp-go/env"
	"lisp-go/eval"
	"lisp-go/lexer"
	"lisp-go/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	environment := env.New()
	scanner := bufio.NewScanner(in)

	if scanner.Err() != nil {
		panic("scanner has error.")
	}

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		objs, err := p.Parse()
		if err != nil {
			fmt.Fprintf(out, "Parse error: %v\n", err)
			continue
		}

		result, err := eval.Eval(objs, environment)
		if err != nil {
			fmt.Fprintf(out, "Eval error: %v\n", err)
			continue
		}
		fmt.Fprintf(out, "[%T] %+v\n", result, result)
	}
}
