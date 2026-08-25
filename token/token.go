package token

var keywords = map[string]TokenType{
	"define": DEFINE,
	"if":     IF,
	"#t":     TRUE,
	"#f":     FALSE,
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

	// operators
	PLUS    = "+"
	MUL     = "*"
	LESS    = "<"
	GREATER = ">"
	EQ      = "="
	NOT_EQ  = "!="

	NE = "!"

	LPAREN = "("
	RPAREN = ")"

	//keywords
	TRUE   = "#t"
	FALSE  = "#f"
	DEFINE = "DEFINE"
	IF     = "IF"
)

func LookupSymbol(symbol string) TokenType {
	if tok, ok := keywords[symbol]; ok {
		return tok
	}

	return SYMBOL
}
