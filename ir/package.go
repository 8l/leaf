package ir

import (
	"e8vm.net/leaf/ir/types"
)

func newPackage(path string) *Package {
	ret := new(Package)
	ret.path = path
	return ret
}

type Package struct {
	path  string // absolute path of the package
	pid   int
	build *Build

	imports  []int // packageIds
	types    []int // map to type id in the package type list
	symbols  []Symbol
	dataSegs []*Data
	codeSegs []*Code
}

func (self *Package) NewFile(name string) *File {
	ret := new(File)
	ret.pack = self
	ret.name = name
	return ret
}

func (self *Package) Save() {
	panic("todo")
}

func (self *Package) DeclType(name string, t types.Type) Symbol {
	panic("todo")
}
