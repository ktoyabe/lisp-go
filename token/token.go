package token

var keywords = map[string]TokenType{
	"define": DEFINE,
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
	PLUS = "+"
	MUL  = "*"

	LPAREN = "("
	RPAREN = ")"

	//keywords
	DEFINE = "DEFINE"
)

func LookupSymbol(symbol string) TokenType {
	if tok, ok := keywords[symbol]; ok {
		return tok
	}

	return SYMBOL
}
