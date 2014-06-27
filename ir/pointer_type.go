package ir

type PointerType struct {
	pack *Package
	t    TypeRef
}

func (self *PointerType) Size() uint32 {
	return addrSize
}

func (self *PointerType) Type() Type {
	return self.pack.Type(self.t)
}
