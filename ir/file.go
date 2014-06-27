package ir

import (
	"e8vm.net/leaf/ir/types"
)

type File struct {
	name    string
	pack    *Package
	imports map[string]int
}

func (self *File) DeclareFunc(name string, t types.Type) *Func {
	ret := new(Func)
	ret.name = name
	ret.t = t
	ret.file = self
	// TODO add symbol into package
	return ret
}

func (self *File) Import(path string) *Package {
	panic("todo")
}
