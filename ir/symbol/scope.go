package symbol

type Scope struct {
	list []Symbol
	m    map[string]Symbol
}

func NewScope() *Scope {
	ret := new(Scope)
	ret.m = make(map[string]Symbol)
	return ret
}

// Try to register a new symbol, returns nil when succeeds,
// returns the symbol with the same name if the symbol is already
// registered
func (self *Scope) Register(name string, c Class, s Symbol) Symbol {
	cur := self.Get(name)
	if cur != nil {
		return cur
	}

	self.list = append(self.list, s)
	self.m[name] = s
	return nil
}

// Scope the symbols in registered order.
// For simplicity, we just returned the list saved in the struct.
// The caller should not change the order of this symbol list.
func (self *Scope) List() []Symbol {
	return self.list
}

// Get the symbol with a specific name
func (self *Scope) Get(n string) Symbol {
	return self.m[n]
}
