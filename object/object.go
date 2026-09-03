package object

import (
	"fmt"
	"strings"
)

func toDebugString(obj Object) string {
	return fmt.Sprintf("[%T] %+v", obj, obj)
}

type Object interface {
	ToString() string
	ToDebugString() string
}

type VoidObject struct {
}

func (o VoidObject) ToString() string {
	return ""
}

func (o VoidObject) ToDebugString() string {
	return toDebugString(o)
}

type IntObject struct {
	Value int
}

func (o *IntObject) ToString() string {
	return fmt.Sprintf("%d", o.Value)
}

func (o *IntObject) ToDebugString() string {
	return toDebugString(o)
}

type FloatObject struct {
	Value float64
}

func (o *FloatObject) ToString() string {
	return fmt.Sprintf("%f", o.Value)
}

func (o *FloatObject) ToDebugString() string {
	return toDebugString(o)
}

type SymbolObject struct {
	Value string
}

func (o *SymbolObject) ToString() string {
	return o.Value
}

func (o *SymbolObject) ToDebugString() string {
	return toDebugString(o)
}

type OperatorObject struct {
	Value string
}

func (o *OperatorObject) ToString() string {
	return o.Value
}

func (o *OperatorObject) ToDebugString() string {
	return toDebugString(o)
}

type BoolObject struct {
	Value bool
}

func (o *BoolObject) ToString() string {
	if o.Value {
		return "#t"
	} else {
		return "#f"
	}
}

func (o *BoolObject) ToDebugString() string {
	return toDebugString(o)
}

type ListObject struct {
	Value []Object
}

func (o *ListObject) ToString() string {
	return listToString(o.Value)
}

func (o *ListObject) ToDebugString() string {
	return toDebugString(o)
}

type ListDataObject struct {
	Value []Object
}

func (o *ListDataObject) ToString() string {
	return listToString(o.Value)
}

func (o *ListDataObject) ToDebugString() string {
	return toDebugString(o)
}

type LambdaObject struct {
	Params []string
	Body   *ListObject
}

func (o *LambdaObject) ToString() string {
	return "LAMBDA"
}

func (o *LambdaObject) ToDebugString() string {
	return toDebugString(o)
}

type StringObject struct {
	Value string
}

func (o *StringObject) ToString() string {
	return o.Value
}

func (o *StringObject) ToDebugString() string {
	return toDebugString(o)
}

func listToString(list []Object) string {
	strs := []string{}
	for _, obj := range list {
		strs = append(strs, obj.ToString())
	}
	return "(" + strings.Join(strs, " ") + ")"
}
