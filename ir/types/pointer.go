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

const addrSize = 4
