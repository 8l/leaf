package types

type Pointer struct {
	Type Type
}

func (self *Pointer) Size() uint32 {
	return U32.Size()
}

func NewPointer(t Type) *Pointer {
	return &Pointer{t}
}
