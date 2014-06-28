package types

// A function type; represents a function signature
type Func struct {
	Args  []Type
	Names []string
	Ret   Type
}

func NewFunc() *Func {
	ret := new(Func)
	ret.Ret = Void // return nothing
	return ret
}

func (self *Func) AddArg(t Type) {
	self.AddNamedArg("_", t)
}

func (self *Func) AddNamedArg(n string, t Type) {
	self.Names = append(self.Names, n)
	self.Args = append(self.Args, t)
}
