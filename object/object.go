package object

import (
	"fmt"
)

func toStringObject(obj Object) string {
	return fmt.Sprintf("[%T] %+v", obj, obj)
}

type Object interface {
	ToString() string
}

type VoidObject struct {
}

func (o VoidObject) ToString() string {
	return toStringObject(o)
}

type IntObject struct {
	Value int
}

func (o *IntObject) ToString() string {
	return toStringObject(o)
}

type SymbolObject struct {
	Value string
}

func (o *SymbolObject) ToString() string {
	return toStringObject(o)
}

type OperatorObject struct {
	Value string
}

func (o *OperatorObject) ToString() string {
	return toStringObject(o)
}

type BoolObject struct {
	Value bool
}

func (o *BoolObject) ToString() string {
	return toStringObject(o)
}

type ListObject struct {
	Value []Object
}

func (o *ListObject) ToString() string {
	return toStringObject(o)
}

type LambdaObject struct {
	Params []string
	Body   *ListObject
}

func (o *LambdaObject) ToString() string {
	return toStringObject(o)
}

type StringObject struct {
	Value string
}

func (o *StringObject) ToString() string {
	return toStringObject(o)
}
