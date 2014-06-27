package ir

import (
	sym "e8vm.net/leaf/ir/symbol"
	"e8vm.net/leaf/ir/types"
)

type File struct {
	name    string
	pack    *Package
	imports map[string]int
}

func (self *File) NewFunc(name string, t *types.Func) *Func {
	ret := new(Func)
	ret.name = name
	ret.t = t
	ret.file = self
	return ret
}

func (self *File) DeclFunc(f *Func) *sym.Symbol {
	return self.pack.declFunc(f)
}

func (self *File) DeclType(name string, t types.Type) *sym.Symbol {
	return self.pack.declType(name, t)
}

func (self *File) Import(path string) *Package {
	panic("todo")
}

func (self *File) newSymTable() *sym.Table {
	ret := sym.NewTable()

	if self.pack.path != "builtin" {
		ret.PushScope(self.builtInScope())
	}
	ret.PushScope(self.packScope())
	// TODO: add imports

	return ret
}

func (self *File) builtInScope() *sym.Scope {
	return self.pack.build.builtIn.scope
}

func (self *File) packScope() *sym.Scope {
	return self.pack.scope
}
