package token

var keywords = map[string]TokenType{
	"define":  DEFINE,
	"if":      IF,
	"lambda":  LAMBDA,
	"list":    LIST,
	"#t":      TRUE,
	"#f":      FALSE,
	"print":   PRINT,
	"inspect": INSPECT,
	"map":     MAP,
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF = "EOF"

	SYMBOL = "SYMBOL"
	INT    = "INT"
	STRING = "STRING"
	FLOAT  = "FLOAT"
	LIST   = "LIST"

	// operators
	PLUS    = "+"
	MINUS   = "-"
	MUL     = "*"
	LESS    = "<"
	GREATER = ">"
	EQ      = "="
	NOT_EQ  = "!="
	AND     = "&"
	OR      = "|"

	NE = "!"

	LPAREN = "("
	RPAREN = ")"

	//keywords
	TRUE   = "#t"
	FALSE  = "#f"
	DEFINE = "DEFINE"
	IF     = "IF"
	LAMBDA = "lambda"

	// built-in function
	PRINT   = "print"
	INSPECT = "inspect"
	MAP     = "map"
)

func LookupSymbol(symbol string) TokenType {
	if tok, ok := keywords[symbol]; ok {
		return tok
	}

	return SYMBOL
}
