package codegen

import (
	"e8vm.net/leaf/lexer"
	"e8vm.net/leaf/symbol"
	"e8vm.net/leaf/token"
)

type namedType struct {
	name string
	pos  *lexer.Token
	typ  typ
}

var _ typ = new(namedType)

func (t *namedType) Size() uint32        { return t.typ.Size() }
func (t *namedType) String() string      { return t.name }
func (t *namedType) Name() string        { return t.name }
func (t *namedType) Kind() symbol.Kind   { return symbol.Type }
func (t *namedType) Token() *lexer.Token { return t.pos }

func declType(tok *lexer.Token) *namedType {
	assert(tok.Token == token.Ident) // must be an indent token
	return &namedType{tok.Lit, tok, nil}
}
