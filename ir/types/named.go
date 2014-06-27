package types

type Named struct {
	Path string
	Name string
	Type Type
}

func (self *Named) Size() uint32 {
	return self.Type.Size()
}
