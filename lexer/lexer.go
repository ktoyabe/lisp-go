package lexer

import (
	"lisp-go/token"
)

type Lexer struct {
	input       string
	position    int  // current position
	readPostion int  // next current-positon
	ch          byte // current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPostion >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPostion]
	}
	l.position = l.readPostion
	l.readPostion += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPostion >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPostion]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		peek := l.peekChar()
		if !isDigit(peek) {
			tok = newToken(token.MINUS, l.ch)
		} else {
			l.readChar() // consume "-"
			tok.Type = token.INT
			tok.Literal = "-" + l.readNumber()
			return tok
		}
	case '*':
		tok = newToken(token.MUL, l.ch)
	case '<':
		tok = newToken(token.LESS, l.ch)
	case '>':
		tok = newToken(token.GREATER, l.ch)
	case '=':
		tok = newToken(token.EQ, l.ch)
	case '!':
		peek := l.peekChar()
		if peek == '=' {
			l.readChar() // read =
			l.readChar() // for next token
			tok.Literal = "!="
			tok.Type = token.NOT_EQ
			return tok
		} else {
			tok = newToken(token.NE, l.ch)
		}
	case '&':
		tok = newToken(token.AND, l.ch)
	case '|':
		tok = newToken(token.OR, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '#':
		tok.Literal = l.readBool()
		tok.Type = token.LookupSymbol(tok.Literal)
		return tok
	case '"':
		tok.Literal = l.readString()
		tok.Type = token.STRING
		return tok
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readSymbol()
			tok.Type = token.LookupSymbol(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == 'r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readSymbol() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	l.readChar() // consume '"'

	position := l.position // start string position: next of '"' position
	for l.ch != '"' {
		l.readChar()
	}
	str := l.input[position:l.position]
	l.readChar() // consume end of '"'

	return str
}

func (l *Lexer) readBool() string {
	position := l.position
	l.readChar() // read #
	l.readChar() // expected t or f
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
