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
	case *object.BoolObject:
		return v, nil
	case *object.ListObject:
		return eval_list(v, env)
	case *object.SymbolObject:
		return eval_symbol(v, env)
	default:
		return nil, fmt.Errorf("Unsupported type. obj=%v, type=%T", obj, obj)
	}
}

func eval_symbol(obj *object.SymbolObject, env *env.Env) (object.Object, error) {
	o, ok := env.Get(obj.Value)
	if !ok {
		return nil, fmt.Errorf("eval_symbol: failed to Env#Get. obj=%v", obj)
	}
	return o, nil
}

func eval_define(obj *object.ListObject, env *env.Env) (object.Object, error) {
	if len(obj.Value) == 0 {
		return nil, fmt.Errorf("eval_define: ListObject length not 3. len=%d", len(obj.Value))
	}

	symbol, ok := (obj.Value[1]).(*object.SymbolObject)
	if !ok {
		return nil, fmt.Errorf("eval_define: obj.Value[1] not SymbolObject. type=%T, v=%v", obj.Value[1], obj.Value[1])
	}

	env.Set(symbol.Value, obj.Value[2])

	return object.VoidObject{}, nil
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
		case "*":
			return eval_binary_op(obj, env)
		case "<":
			return eval_binary_op(obj, env)
		case ">":
			return eval_binary_op(obj, env)
		case "define":
			return eval_define(obj, env)
		case "if":
			return eval_if(obj, env)
		default:
			return nil, fmt.Errorf("eval_list: Unsupported operator type. head=%v", head)
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

	lhs, err := eval_obj(obj.Value[1], env)
	if err != nil {
		return nil, fmt.Errorf("eval_binary_op: failed to eval_obj(obj.Value[1]). error=%v", err)
	}
	lhs_int, lhs_ok := lhs.(*object.IntObject)

	rhs, err := eval_obj(obj.Value[2], env)
	if err != nil {
		return nil, fmt.Errorf("eval_binary_op: failed to eval_obj(obj.Value[2]). error=%v", err)
	}
	rhs_int, rhs_ok := rhs.(*object.IntObject)

	if !(lhs_ok && rhs_ok) {
		return nil, fmt.Errorf("eval_binary_op support only (op IntObject IntObject)")
	}

	switch op.Value {
	case "+":
		return &object.IntObject{Value: lhs_int.Value + rhs_int.Value}, nil
	case "*":
		return &object.IntObject{Value: lhs_int.Value * rhs_int.Value}, nil
	case "<":
		return &object.BoolObject{Value: lhs_int.Value < rhs_int.Value}, nil
	case ">":
		return &object.BoolObject{Value: lhs_int.Value > rhs_int.Value}, nil
	default:
		return nil, fmt.Errorf("eval_binary_op unsuppoted operator. op=%v", op)
	}
}

func eval_if(obj *object.ListObject, env *env.Env) (object.Object, error) {
	if len(obj.Value) != 4 {
		return nil, fmt.Errorf("eval_if: size must be 4. length=%d", len(obj.Value))
	}
	cond, err := eval_obj(obj.Value[1], env)
	if err != nil {
		return nil, fmt.Errorf("eval_if: failed to eval_obj(cond_cell). error=%v", err)
	}
	result, ok := cond.(*object.BoolObject)
	if !ok {
		return nil, fmt.Errorf("eval_if: eval_obj(cond) not BoolObject. cond=%v, type=%v", cond, cond)
	}
	if result.Value {
		return eval_obj(obj.Value[2], env)
	} else {
		return eval_obj(obj.Value[3], env)
	}
}
