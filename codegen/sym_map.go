package codegen

import (
	"fmt"
	"reflect"
)

type symMap struct {
	tab map[string]symbol
}

func newSymMap() *symMap {
	ret := new(symMap)
	ret.tab = make(map[string]symbol)
	return ret
}

func symName(sym symbol) string {
	switch sym := sym.(type) {
	case *namedType:
		return sym.String()
	case *function:
		return sym.Name
	}

	panic(fmt.Sprintf("not a symbol: %s", reflect.TypeOf(sym).Name))
}

func (self *symMap) add(name string, sym symbol) {
	self.tab[name] = sym
}

func (self *symMap) Add(syms ...symbol) {
	for _, sym := range syms {
		name := symName(sym)
		self.add(name, sym)
	}
}

func (self *symMap) TryAdd(sym symbol) symbol {
	name := symName(sym)
	s := self.Get(name)
	if s != nil {
		return s
	}
	self.add(name, sym)
	return nil
}

func (self *symMap) Get(name string) symbol {
	return self.tab[name]
}
