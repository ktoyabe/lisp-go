package eval

import (
	"fmt"
	"lisp-go/env"
	"lisp-go/object"
)

func Eval(obj object.Object, env *env.Env) (object.Object, error) {
	return evalObj(obj, env)
}

func evalObj(obj object.Object, env *env.Env) (object.Object, error) {
	switch v := obj.(type) {
	case *object.IntObject:
		return v, nil
	case *object.BoolObject:
		return v, nil
	case *object.StringObject:
		return v, nil
	case *object.FloatObject:
		return v, nil
	case *object.ListObject:
		return evalList(v, env)
	case *object.SymbolObject:
		return evalSymbol(v, env)
	default:
		return nil, fmt.Errorf("evalObj: Unsupported type. obj=%v, type=%T", obj, obj)
	}
}

func evalSymbol(obj *object.SymbolObject, env *env.Env) (object.Object, error) {
	o, ok := env.Get(obj.Value)
	if !ok {
		return nil, fmt.Errorf("eval_symbol: Unknown symbol name. name=\"%v\"", obj.Value)
	}
	return o, nil
}

func evalDefine(obj *object.ListObject, env *env.Env) (object.Object, error) {
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

func evalFunctionCall(lambda *object.LambdaObject, args []object.Object, environment *env.Env) (object.Object, error) {
	if len(lambda.Params) != len(args) {
		return nil, fmt.Errorf("eval_function_call: Params length error. len(Lambda#Params)=%d, actual=%d", len(lambda.Params), len(args))
	}

	functionScopeEnv := env.Extend(environment)
	for i, p := range lambda.Params {
		v, err := evalObj(args[i], environment)
		if err != nil {
			return nil, err
		}
		functionScopeEnv.Set(p, v)
	}

	return evalList(lambda.Body, functionScopeEnv)
}

func evalList(obj *object.ListObject, env *env.Env) (object.Object, error) {
	if len(obj.Value) == 0 {
		return nil, fmt.Errorf("ListObject length not zero.")
	}

	head := obj.Value[0]

	switch v := head.(type) {
	case *object.OperatorObject:
		if len(obj.Value) != 3 { // (operator lhs rhs)
			return nil, fmt.Errorf("eval_list: operator object require 3 elements. got=%d", len(obj.Value))
		}
		return evalBinaryOp(v, obj.Value[1], obj.Value[2], env)
	case *object.SymbolObject:
		switch v.Value {
		case "define":
			return evalDefine(obj, env)
		case "if":
			return evalIf(obj, env)
		default:
			v, err := evalSymbol(v, env)
			if err != nil {
				return nil, err
			}
			switch vv := v.(type) {
			case *object.LambdaObject:
				return evalFunctionCall(vv, obj.Value[1:], env)
			default:
				return nil, fmt.Errorf("eval_list: Unsupported operator type. head=%v", head)
			}
		}
	default:
		return nil, fmt.Errorf("Unsupported type. head=%v, type=%T", head, head)
	}
}

func evalBinaryOpInt(op *object.OperatorObject, lhs_int *object.IntObject, rhs_int *object.IntObject) (object.Object, error) {
	switch op.Value {
	case "+":
		return &object.IntObject{Value: lhs_int.Value + rhs_int.Value}, nil
	case "-":
		return &object.IntObject{Value: lhs_int.Value - rhs_int.Value}, nil
	case "*":
		return &object.IntObject{Value: lhs_int.Value * rhs_int.Value}, nil
	case "<":
		return &object.BoolObject{Value: lhs_int.Value < rhs_int.Value}, nil
	case ">":
		return &object.BoolObject{Value: lhs_int.Value > rhs_int.Value}, nil
	case "=":
		return &object.BoolObject{Value: lhs_int.Value == rhs_int.Value}, nil
	case "!=":
		return &object.BoolObject{Value: lhs_int.Value != rhs_int.Value}, nil
	default:
		return nil, fmt.Errorf("evalBinaryOpInt: unsuppoted operator. op=%v", op)
	}
}

func evalBinaryOpBool(op *object.OperatorObject, lhs *object.BoolObject, rhs *object.BoolObject) (object.Object, error) {
	switch op.Value {
	case "&":
		return &object.BoolObject{Value: lhs.Value && rhs.Value}, nil
	case "|":
		return &object.BoolObject{Value: lhs.Value || rhs.Value}, nil
	default:
		return nil, fmt.Errorf("evalBinaryOpBool: unsupported operator. op=%v", op)
	}
}

func evalBinaryOpString(op *object.OperatorObject, lhs *object.StringObject, rhs *object.StringObject) (object.Object, error) {
	switch op.Value {
	case "=":
		return &object.BoolObject{Value: lhs.Value == rhs.Value}, nil
	case "!=":
		return &object.BoolObject{Value: lhs.Value != rhs.Value}, nil
	case "+":
		return &object.StringObject{Value: lhs.Value + rhs.Value}, nil
	default:
		return nil, fmt.Errorf("evalBinaryOpString: unsupported operator. op=%v", op)
	}
}

func evalBinaryOpFloat(op *object.OperatorObject, lhs float64, rhs float64) (object.Object, error) {
	switch op.Value {
	case "+":
		return &object.FloatObject{Value: lhs + rhs}, nil
	default:
		return nil, fmt.Errorf("evalBinaryOpFloat: unsupported operator=%v", op)
	}
}

func evalBinaryOp(op *object.OperatorObject, lhs object.Object, rhs object.Object, env *env.Env) (object.Object, error) {
	lhs_obj, err := evalObj(lhs, env)
	if err != nil {
		return nil, fmt.Errorf("eval_binary_op: failed to eval_obj(lhs). error=%v", err)
	}
	rhs_obj, err := evalObj(rhs, env)
	if err != nil {
		return nil, fmt.Errorf("eval_binary_op: failed to eval_obj(rhs). error=%v", err)
	}
	switch l := lhs_obj.(type) {
	case *object.IntObject:
		switch r := rhs_obj.(type) {
		case *object.IntObject:
			return evalBinaryOpInt(op, l, r)
		default:
			return nil, fmt.Errorf("evalBinaryOp: unsupported (op, lhs, rhs) type. op=%v, lhsType=%T, rhsType=%T", op, lhs_obj, rhs_obj)
		}
	case *object.FloatObject:
		switch r := rhs_obj.(type) {
		case *object.FloatObject:
			return evalBinaryOpFloat(op, l.Value, r.Value)
		default:
			return nil, fmt.Errorf("evalBinaryOp: unsupported (op, lhs, rhs) type. op=%v, lhsType=%T, rhsType=%T", op, lhs_obj, rhs_obj)
		}
	case *object.BoolObject:
		switch r := rhs_obj.(type) {
		case *object.BoolObject:
			return evalBinaryOpBool(op, l, r)
		default:
			return nil, fmt.Errorf("evalBinaryOp: unsupported (op, lhs, rhs) type. op=%v, lhsType=%T, rhsType=%T", op, lhs_obj, rhs_obj)
		}
	case *object.StringObject:
		switch r := rhs_obj.(type) {
		case *object.StringObject:
			return evalBinaryOpString(op, l, r)
		default:
			return nil, fmt.Errorf("evalBinaryOp: unsupported (op, lhs, rhs) type. op=%v, lhsType=%T, rhsType=%T", op, lhs_obj, rhs_obj)
		}
	default:
		return nil, fmt.Errorf("evalBinaryOp: unsupported (op, lhs, rhs) type. op=%v, lhsType=%T, rhsType=%T", op, lhs_obj, rhs_obj)
	}
}

func evalIf(obj *object.ListObject, env *env.Env) (object.Object, error) {
	if len(obj.Value) != 4 {
		return nil, fmt.Errorf("eval_if: size must be 4. length=%d", len(obj.Value))
	}
	cond, err := evalObj(obj.Value[1], env)
	if err != nil {
		return nil, fmt.Errorf("eval_if: failed to eval_obj(cond_cell). error=%v", err)
	}
	result, ok := cond.(*object.BoolObject)
	if !ok {
		return nil, fmt.Errorf("eval_if: eval_obj(cond) not BoolObject. cond=%v, type=%v", cond, cond)
	}
	if result.Value {
		return evalObj(obj.Value[2], env)
	} else {
		return evalObj(obj.Value[3], env)
	}
}
