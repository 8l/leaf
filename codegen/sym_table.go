package codegen

type symTable struct {
	builtIn *symMap
	tops    *symMap

	// we will swap this with different file import symbols for
	// different files
	// when building the imports for each file,
	// we also need to check that thes imports does not collide
	// with top decls
	imports *symMap

	scopes []*symMap
}

func newSymTable() *symTable {
	ret := new(symTable)
	ret.builtIn = builtIn
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

// Returns nil on succeed; return the symbol entry if it is already defined
func (self *symTable) Define(s symbol) *symEntry {
	top := self.top()
	return top.TryAdd(s)
}

func (self *symTable) DeclTop(name string, kind symKind) *symEntry {
	return self.tops.TryDecl(name, kind)
}

func (self *symTable) Import() symbol {
	panic("todo")
}

// Search in the current scope hierarchy for a symbol name
func (self *symTable) Find(name string) *symEntry {
	nscope := len(self.scopes)

	// look in the scopes
	for i := nscope - 1; i >= 0; i-- {
		s := self.scopes[i].Get(name)
		if s != nil {
			return s
		}
	}

	if self.imports != nil {
		// check the imports
		if s := self.imports.Get(name); s != nil {
			return s
		}
	}

	// top level symbols?
	if s := self.tops.Get(name); s != nil {
		return s
	}
	// finally we check built-in
	if s := self.builtIn.Get(name); s != nil {
		return s
	}

	return nil
}
