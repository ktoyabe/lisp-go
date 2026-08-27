package object

import (
	"strconv"
)

type Object interface {
	TokenLiteral() string
}

type VoidObject struct {
	Object
}

func (o VoidObject) TokenLiteral() string {
	return "Void"
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

type BoolObject struct {
	Object
	Value bool
}

func (o *BoolObject) TokenLiteral() string {
	if o.Value {
		return "#t"
	} else {
		return "#f"
	}
}

type ListObject struct {
	Object
	Value []Object
}

func (o *ListObject) TokenLiteral() string {
	return "List"
}

type LambdaObject struct {
	Object
	Params []string
	Body   *ListObject
}

func (o *LambdaObject) TokenLiteral() string {
	return "Lambda"
}
