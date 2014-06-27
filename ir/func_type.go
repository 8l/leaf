package ir

type FuncType struct {
	pack *Package
	args []TypeRef
	ret  TypeRef
}

func (self *FuncType) AddArg(t TypeRef) {
	self.args = append(self.args, t)
}

func (self *FuncType) SetRet(t TypeRef) {
	self.ret = t
}

func (self *FuncType) Size() uint32 {
	return addrSize
}

func (self *FuncType) RetType() Type {
	return self.pack.Type(self.ret)
}

func (self *FuncType) ArgType(i int) Type {
	return self.pack.Type(self.args[i])
}
