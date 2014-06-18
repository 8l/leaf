package codegen

type Func struct {
	Name string
	Args []*Arg
	Ret  Type
	// Impl // implementation
}

type Arg struct {
	Name string
	Type Type
}

func NewFunc(name string) *Func {
	ret := new(Func)
	ret.Name = name
	return ret
}

func (f *Func) AddArg(t Type) {
	f.Args = append(f.Args, &Arg{Type: t})
}

func (f *Func) AddNamedArg(t Type, name string) {
	f.Args = append(f.Args, &Arg{Name: name, Type: t})
}
