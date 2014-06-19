package codegen

type function struct {
	name string
	args []*funcArg
	ret  typ
}

type funcArg struct {
	name string
	typ  typ
}

func newFunc(name string) *function {
	ret := new(function)
	ret.name = name
	ret.ret = tpVoid
	return ret
}

func (f *function) addArg(t typ) {
	f.args = append(f.args, &funcArg{typ: t})
}

func (f *function) addNamedArg(t typ, name string) {
	f.args = append(f.args, &funcArg{name: name, typ: t})
}
