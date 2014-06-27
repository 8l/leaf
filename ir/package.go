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

	files    []*File
	dataSegs []*Data
	codeSegs []*Code
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

func (self *Package) DeclType(name string, t types.Type) *sym.Symbol {
	return self.scope.Add(name, sym.Type, t)
}

func (self *Package) DeclFunc(f *Func) *sym.Symbol {
	return self.scope.Add(f.name, sym.Func, f)
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

	fn := f.NewFunc("printInt", types.NewFunc(nil, types.Int32))
	f.DeclFunc(fn)

	return p
}
