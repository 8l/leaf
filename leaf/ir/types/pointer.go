package types

type Pointer struct {
	Type Type
}

func NewPointer(t Type) *Pointer {
	ret := new(Pointer)
	ret.Type = t
	return ret
}

func (self *Pointer) Size() uint32 {
	return addrSize
}

func (self *Pointer) String() string {
	return "*" + self.Type.String()
}

const addrSize = 4
