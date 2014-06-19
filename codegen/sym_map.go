package codegen

type symMap struct {
	tab map[string]*symEntry
}

type symEntry struct {
	name string
	kind symKind
	s    symbol
}

func newSymMap() *symMap {
	ret := new(symMap)
	ret.tab = make(map[string]*symEntry)
	return ret
}

func getSymKind(sym symbol) symKind {
	switch sym.(type) {
	case *namedType:
		return symType
	case *function:
		return symFunc
	}

	panic("not a sym")
}

func (self *symMap) add(name string, kind symKind, sym symbol) {
	self.tab[name] = &symEntry{name, kind, sym}
}

func (self *symMap) decl(name string, kind symKind) {
	self.add(name, kind, nil)
}

func (self *symMap) Add(syms ...symbol) {
	for _, sym := range syms {
		name := sym.Name()
		kind := getSymKind(sym)
		self.add(name, kind, sym)
	}
}

func (self *symMap) TryAdd(sym symbol) *symEntry {
	name := sym.Name()
	s := self.Get(name)
	if s != nil {
		return s
	}
	self.add(name, getSymKind(sym), sym)
	return nil
}

func (self *symMap) TryDecl(name string, kind symKind) *symEntry {
	s := self.Get(name)
	if s != nil {
		return s
	}
	self.decl(name, kind)
	return nil
}

func (self *symMap) Get(name string) *symEntry {
	return self.tab[name]
}
