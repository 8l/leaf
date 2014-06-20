package types

type Pointer struct {
	Type Type
}

func (self *Pointer) Size() uint32 {
	return _uint32.Size()
}

func NewPointer(t Type) *Pointer {
	return &Pointer{t}
}
