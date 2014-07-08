package symbol

type Scope struct {
	list []*Symbol
	m    map[string]*Symbol
}

func NewScope() *Scope {
	ret := new(Scope)
	ret.m = make(map[string]*Symbol)
	return ret
}

func (self *Scope) Add(name string, c Class, obj interface{}) *Symbol {
	cur := self.Get(name)
	if cur != nil {
		panic("symbol exists")
	}

	sym := new(Symbol)
	sym.name = name
	sym.class = c
	sym.object = obj

	self.list = append(self.list, sym)
	self.m[name] = sym

	return sym
}

func (self *Scope) List() []*Symbol {
	return self.list
}

func (self *Scope) Get(n string) *Symbol {
	return self.m[n]
}

func (self *Scope) Has(n string) bool {
	return self.Get(n) != nil
}
