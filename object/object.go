package object

import (
	"strconv"
)

type Object interface {
	TokenLiteral() string
}

type VoidObject struct {
}

func (o VoidObject) TokenLiteral() string {
	return "Void"
}

type IntObject struct {
	Value int
}

func (o *IntObject) TokenLiteral() string {
	return strconv.Itoa(o.Value)
}

type SymbolObject struct {
	Value string
}

func (o *SymbolObject) TokenLiteral() string {
	return o.Value
}

type OperatorObject struct {
	Value string
}

func (o *OperatorObject) TokenLiteral() string {
	return o.Value
}

type BoolObject struct {
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
	Value []Object
}

func (o *ListObject) TokenLiteral() string {
	return "List"
}

type LambdaObject struct {
	Params []string
	Body   *ListObject
}

func (o *LambdaObject) TokenLiteral() string {
	return "Lambda"
}
