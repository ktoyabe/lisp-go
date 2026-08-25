package object

import "strconv"

type Object interface {
	TokenLiteral() string
}

type IntObject struct {
	Object
	Value int
}

func (o *IntObject) TokenLiteral() string {
	return strconv.Itoa(o.Value)
}

type SymbolObject struct {
	Object
	Value string
}

func (o *SymbolObject) TokenLiteral() string {
	return o.Value
}

type ListObject struct {
	Object
	Value []Object
}

func (o *ListObject) TokenLiteral() string {
	return "List"
}
