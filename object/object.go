package object

import (
	"fmt"
	"io"
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
	return fmt.Sprintf("[%T] &{Value:[%s]}", o, _listToInspectString(o.Value, " "))
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
	return "(" + _listToString(list, " ") + ")"
}

func _listToString(list []Object, delim string) string {
	strs := []string{}
	for _, obj := range list {
		strs = append(strs, obj.ToString())
	}
	return strings.Join(strs, delim)
}

func _listToInspectString(list []Object, delim string) string {
	strs := []string{}
	for _, obj := range list {
		strs = append(strs, obj.ToDebugString())
	}
	return strings.Join(strs, delim)
}

type ObjectPrinter func(obj Object, out io.Writer) (int, error)

func PrintObject(obj Object, out io.Writer) (int, error) {
	str := obj.ToString()
	if str != "" {
		return fmt.Fprintf(out, "%s\n", str)
	} else {
		return 0, nil
	}
}

func InspectObject(obj Object, out io.Writer) (int, error) {
	return fmt.Fprintf(out, "%v\n", obj.ToDebugString())
}
