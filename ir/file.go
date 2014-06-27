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
	return self.pack.DeclFunc(f)
}

func (self *File) DeclType(name string, t types.Type) *sym.Symbol {
	return self.pack.DeclType(name, t)
}

func (self *File) Import(path string) *Package {
	panic("todo")
}

func (self *File) newSymTable() *sym.Table {
	ret := sym.NewTable()
	ret.PushScope(self.builtInScope())
	return ret
}

func (self *File) builtInScope() *sym.Scope {
	return self.pack.build.builtIn.scope
}
