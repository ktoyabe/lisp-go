package main

import (
	"lisp-go/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
