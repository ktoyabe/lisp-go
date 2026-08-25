package env

import "lisp-go/object"

type Env struct {
	parent *Env
	vars   map[string]object.Object
}

func New() *Env {
	return &Env{vars: make(map[string]object.Object)}
}

func Extend(parent *Env) *Env {
	return &Env{parent: parent, vars: make(map[string]object.Object)}
}

func (env *Env) Get(name string) (object.Object, bool) {
	val, ok := env.vars[name]
	if ok {
		return val, true
	}
	if env.parent == nil {
		return nil, false
	}
	return env.parent.Get(name)
}

func (env *Env) Set(name string, value object.Object) {
	env.vars[name] = value
}
