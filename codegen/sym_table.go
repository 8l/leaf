package codegen

type symTable struct {
	builtin *symMap
	imports *symMap
	tops    *symMap

	scopes []*symMap
}

func newSymTable() *symTable {
	ret := new(symTable)
	ret.builtin = makeBuiltIn()
	ret.imports = newSymMap()
	ret.tops = newSymMap()

	return ret
}

func (self *symTable) EnterScope() {
	self.scopes = append(self.scopes, newSymMap())
}

func (self *symTable) ExitScope() {
	nscope := len(self.scopes)
	assert(nscope > 0)
	self.scopes = self.scopes[:nscope-1]
}

func (self *symTable) top() *symMap {
	nscope := len(self.scopes)
	if nscope == 0 {
		return self.tops // top declares
	}

	return self.scopes[nscope-1]
}

// Returns nil on succeed; return the symbol if it is already defined
func (self *symTable) Define(s symbol) symbol {
	top := self.top()
	return top.TryAdd(s)
}

// Search in the current scope hierarchy for a symbol name
func (self *symTable) Find(name string) symbol {
	nscope := len(self.scopes)

	for i := nscope - 1; i >= 0; i-- {
		s := self.scopes[i].Get(name)
		if s != nil {
			return s
		}
	}

	if s := self.tops.Get(name); s != nil {
		return s
	}
	if s := self.imports.Get(name); s != nil {
		return s
	}
	if s := self.builtin.Get(name); s != nil {
		return s
	}

	return nil
}
