package codegen

type function struct {
	Name string
	Args []*funcArg
	Ret  Type
	// Impl // implementation
}

type funcArg struct {
	Name string
	Type Type
}

func newFunc(name string) *function {
	ret := new(function)
	ret.Name = name
	return ret
}

func (f *function) AddArg(t Type) {
	f.Args = append(f.Args, &funcArg{Type: t})
}

func (f *function) AddNamedArg(t Type, name string) {
	f.Args = append(f.Args, &funcArg{Name: name, Type: t})
}
