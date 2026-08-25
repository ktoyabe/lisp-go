package eval

import (
	"fmt"
	"lisp-go/env"
	"lisp-go/object"
)

func Eval(obj object.Object, env *env.Env) (object.Object, error) {
	return eval_obj(obj, env)
}

func eval_obj(obj object.Object, env *env.Env) (object.Object, error) {
	switch v := obj.(type) {
	case *object.IntObject:
		return v, nil
	case *object.ListObject:
		return eval_list(v, env)
	default:
		return nil, fmt.Errorf("Unsupported type. obj=%v, type=%T", obj, obj)
	}
}

func eval_list(obj *object.ListObject, env *env.Env) (object.Object, error) {
	if len(obj.Value) == 0 {
		return nil, fmt.Errorf("ListObject length not zero.")
	}

	head := obj.Value[0]
	switch v := head.(type) {
	case *object.SymbolObject:
		switch v.Value {
		case "+":
			return eval_binary_op(obj, env)
		default:
			return nil, fmt.Errorf("Unsupported operator type. head=%v", head)
		}
	default:
		return nil, fmt.Errorf("Unsupported type. head=%v, type=%T", head, head)
	}
}

func eval_binary_op(obj *object.ListObject, env *env.Env) (object.Object, error) {
	if len(obj.Value) != 3 {
		return nil, fmt.Errorf("binary operator size must be 3. length=%d", len(obj.Value))
	}
	head := obj.Value[0]
	op, ok := head.(*object.SymbolObject)
	if !ok {
		return nil, fmt.Errorf("eval_binary_op head object not SymbolObject. head=%v, type=%T", head, head)
	}

	lhs := obj.Value[1]
	lhs_int, lhs_ok := lhs.(*object.IntObject)

	rhs := obj.Value[2]
	rhs_int, rhs_ok := rhs.(*object.IntObject)

	if !(lhs_ok && rhs_ok) {
		return nil, fmt.Errorf("eval_binary_op support only (op IntObject IntObject)")
	}

	switch op.Value {
	case "+":
		return &object.IntObject{Value: lhs_int.Value + rhs_int.Value}, nil
	default:
		return nil, fmt.Errorf("eval_binary_op unsuppoted operator. op=%v", op)
	}
}
