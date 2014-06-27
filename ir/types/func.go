package types

type Func struct {
	Ret  Type
	Args []Type
}

func (self *Func) Size() uint32 {
	return addrSize
}

func (self *Func) AddArg(t Type) {
	self.Args = append(self.Args, t)
}
