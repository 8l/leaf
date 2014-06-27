package ir

import (
	"e8vm.net/leaf/ir/types"
	"e8vm.net/leaf/ir/symbol"
)

func newPackage(path string) *Package {
	ret := new(Package)
	ret.path = path
	ret.scope = symbol.NewScope()

	return ret
}

type Package struct {
	path  string // absolute path of the package
	pid   int
	build *Build
	scope *symbol.Scope
	
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

func (self *Package) DeclType(name string, t types.Type) symbol.Symbol {
	return self.scope.Register(name, symbol.Type, t)
}
