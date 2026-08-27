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
			fmt.Printf("Parse error: %v\n", err)
			continue
		}

		result, err := eval.Eval(objs, environment)
		if err != nil {
			fmt.Printf("Eval error: %v\n", err)
			continue
		}
		fmt.Printf("[%T] %+v\n", result, result)
	}
}
