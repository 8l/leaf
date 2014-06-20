package types

import (
	"e8vm.net/leaf/lexer"
	"e8vm.net/leaf/symbol"
)

// Named type serves as an alias for a type.
// It can be used as a symbol and be put in a symbol table,
// and can be resolved later.
type Named struct {
	name  string
	token *lexer.Token

	Type Type
}

var _ symbol.Symbol = new(Named)

func NewNamed(n string, tok *lexer.Token) *Named {
	ret := new(Named)
	ret.name = n
	ret.token = tok
	return ret
}

func newBuiltIn(n string, t Type) *Named {
	ret := new(Named)
	ret.name = n
	ret.Type = t
	return ret
}

// Size of the type, will panic when it is not resolved
func (self *Named) Size() uint32 {
	return self.Type.Size()
}

func (self *Named) String() string {
	return self.name
}

func (self *Named) Name() string {
	return self.name
}

func (self *Named) Kind() symbol.Kind {
	return symbol.Type
}

func (self *Named) Token() *lexer.Token {
	return self.token
}
