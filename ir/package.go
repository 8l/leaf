package ir

import (
	sym "e8vm.net/leaf/ir/symbol"
	"e8vm.net/leaf/ir/types"
)

func newPackage(path string) *Package {
	ret := new(Package)
	ret.path = path
	ret.scope = sym.NewScope()

	return ret
}

type Package struct {
	path  string     // absolute path of the package
	build *Build     // the build that this package belongs to
	scope *sym.Scope // top level symbols

	files []*File
	types []*sym.Symbol
	funcs []*sym.Symbol
}

func (self *Package) NewFile(name string) *File {
	ret := new(File)
	ret.pack = self
	ret.name = name

	self.files = append(self.files, ret)
	return ret
}

func (self *Package) Save() {
	panic("todo")
}

func (self *Package) declType(name string, t types.Type) *sym.Symbol {
	ret := self.scope.Add(name, sym.Type, t)
	self.types = append(self.types, ret)
	return ret
}

func (self *Package) declFunc(f *Func) *sym.Symbol {
	ret := self.scope.Add(f.name, sym.Func, f)
	self.funcs = append(self.funcs, ret)
	return ret
}

func makeBuiltIn() *Package {
	p := newPackage("builtin")
	f := p.NewFile("builtin.leaf")

	f.DeclType("int32", types.Int32)
	f.DeclType("uint32", types.Uint32)
	f.DeclType("int", types.Int32)
	f.DeclType("uint", types.Uint32)
	f.DeclType("int8", types.Int8)
	f.DeclType("uint8", types.Uint8)
	f.DeclType("char", types.Int8)
	f.DeclType("byte", types.Uint8)
	f.DeclType("ptr", types.NewPointer(nil))

	f.DeclFunc(f.NewFunc("printInt", types.NewFunc(nil, types.Int32)))
	f.DeclFunc(f.NewFunc("printChar", types.NewFunc(nil, types.Int8)))

	return p
}
