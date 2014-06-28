package codegen

import (
	"e8vm.net/leaf/codegen/symbol"
	"e8vm.net/leaf/codegen/types"
	"e8vm.net/leaf/lexer"
)

type function struct {
	*types.Func

	name  string
	token *lexer.Token
	addr  uint32 // this will be filled later
}

type funcArg struct {
	name string
	typ  types.Type
}

func newFunc(t *lexer.Token) *function {
	ret := new(function)
	ret.name = t.Lit
	ret.token = t
	return ret
}

func newBuiltInFunc(name string) *function {
	ret := new(function)
	ret.name = name
	return ret
}

func (f *function) Name() string        { return f.name }
func (f *function) Kind() symbol.Kind   { return symbol.Func }
func (f *function) Token() *lexer.Token { return f.token }
