package codegen

import (
	"fmt"
	"reflect"
)

type Sym interface{}

type SymMap struct {
	tab map[string]Sym
}

func newSymMap() *SymMap {
	ret := new(SymMap)
	ret.tab = make(map[string]Sym)
	return ret
}

func symName(sym Sym) string {
	switch sym := sym.(type) {
	case *BasicType:
		return sym.Name
	case *NamedType:
		return sym.Name
	case *Func:
		return sym.Name
	}

	panic(fmt.Sprintf("bug sym type: %s", reflect.TypeOf(sym).Name))
}

func (self *SymMap) add(name string, sym Sym) {
	self.tab[name] = sym
}

func (self *SymMap) Add(syms ...Sym) {
	for _, sym := range syms {
		name := symName(sym)
		self.add(name, sym)
	}
}

func (self *SymMap) TryAdd(sym Sym) Sym {
	name := symName(sym)
	s := self.Get(name)
	if s != nil {
		return s
	}
	self.add(name, sym)
	return nil
}

func (self *SymMap) Get(name string) Sym {
	return self.tab[name]
}
