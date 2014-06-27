package ir

type File struct {
	pack    *Package
	imports map[string]int
}

func (self *File) DeclareFunc(name string, t TypeRef) *Func {
	ret := new(Func)
	ret.name = name
	ret.t = t
	ret.file = self
	// TODO add symbol into package
	return ret
}
