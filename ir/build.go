package ir

import (
	"e8vm.net/leaf/ir/types"
)

type Build struct {
	packs   []*Package
	packMap map[string]*Package
}

func NewBuild() *Build {
	ret := new(Build)
	ret.packMap = make(map[string]*Package)

	ret.addBuiltIn()

	return ret
}

func (self *Build) addBuiltIn() {
	p := newPackage("<builtin>")
	self.addPackage(p)

	p.DeclType("int32", types.Int32)
	p.DeclType("uint32", types.Uint32)
	p.DeclType("int", types.Int32)
	p.DeclType("uint", types.Uint32)
	p.DeclType("int8", types.Int8)
	p.DeclType("uint8", types.Uint8)
	p.DeclType("char", types.Int8)
	p.DeclType("byte", types.Uint8)
	p.DeclType("ptr", types.NewPointer(nil))
}

func (self *Build) addPackage(p *Package) {
	if self.packMap[p.path] != nil {
		panic("package with the same path already exists")
	}

	p.pid = len(self.packs)
	p.build = self
	self.packs = append(self.packs, p)
	self.packMap[p.path] = p
}

func (self *Build) NewPackage(path string) *Package {
	ret := new(Package)
	ret.path = path
	self.addPackage(ret)

	return ret
}

func (self *Build) ImportPackage(path string) *Package {
	if self.packMap[path] == nil {
		self.importPackage(path)
	}
	return self.packMap[path]
}

// import the header (consts, types and symbols) only
func (self *Build) importPackage(path string) *Package {
	panic("todo")
}

func (self *Build) LoadPackage(path string) *Package {
	panic("todo")
}

func (self *Func) Define() *Code {
	ret := new(Code)
	self.code = ret
	return ret
}
