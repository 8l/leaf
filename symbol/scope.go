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
func (self *Scope) Register(s Symbol) Symbol {
	name := s.Name()
	cur := self.Get(name)
	if cur != nil {
		return cur
	}

	self.list = append(self.list, s)
	self.m[name] = s
	return nil
}

// Scope the symbols in registered order.
func (self *Scope) Scope() []Symbol {
	ret := make([]Symbol, len(self.list))
	copy(ret, self.list)
	return ret
}

// Scope the symbols of a kind in registred order
func (self *Scope) ScopeKind(k Kind) []Symbol {
	var ret []Symbol
	for _, s := range self.list {
		if s.Kind() == k {
			ret = append(ret, s)
		}
	}

	return ret
}

// Get the symbol with a specific name
func (self *Scope) Get(n string) Symbol {
	return self.m[n]
}
