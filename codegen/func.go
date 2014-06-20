package codegen

import (
	"e8vm.net/leaf/lexer"
	"e8vm.net/leaf/symbol"
	"e8vm.net/leaf/types"
)

type function struct {
	name  string
	args  []*funcArg
	ret   types.Type
	token *lexer.Token
}

type funcArg struct {
	name string
	typ  types.Type
}

func newFunc(t *lexer.Token) *function {
	ret := new(function)
	ret.name = t.Lit
	ret.token = t
	ret.ret = types.Void
	return ret
}

func newBuiltInFunc(name string) *function {
	ret := new(function)
	ret.name = name
	ret.ret = types.Void
	return ret
}

func (f *function) addArg(t types.Type) {
	f.args = append(f.args, &funcArg{typ: t})
}

func (f *function) addNamedArg(t types.Type, name string) {
	f.args = append(f.args, &funcArg{name: name, typ: t})
}

func (f *function) Name() string        { return f.name }
func (f *function) Kind() symbol.Kind   { return symbol.Func }
func (f *function) Token() *lexer.Token { return f.token }
