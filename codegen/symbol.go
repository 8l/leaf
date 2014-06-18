package codegen

type Sym interface{}

type SymTable struct {
	tab map[string]Sym
}

func newSymTable() *SymTable {
	ret := new(SymTable)
	ret.tab = make(map[string]Sym)
	return ret
}

func (self *SymTable) Add(syms ...Sym) {
	for _, sym := range syms {
		switch sym := sym.(type) {
		case *BasicType:
			self.tab[sym.Name] = sym
		case *Func:
			self.tab[sym.Name] = sym
		default:
			panic("bug")
		}
	}
}
