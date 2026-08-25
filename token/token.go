package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF = "EOF"

	SYMBOL = "SYMBOL"
	INT    = "INT"

	PLUS = "+"

	LPAREN = "("
	RPAREN = ")"
)
