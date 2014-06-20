package codegen

import (
	"e8vm.net/leaf/lexer"
	"e8vm.net/leaf/symbol"
)

type function struct {
	name  string
	args  []*funcArg
	ret   typ
	token *lexer.Token
}

type funcArg struct {
	name string
	typ  typ
}

func newFunc(name string) *function {
	ret := new(function)
	ret.name = name
	ret.ret = tpVoid
	return ret
}

func (f *function) addArg(t typ) {
	f.args = append(f.args, &funcArg{typ: t})
}

func (f *function) addNamedArg(t typ, name string) {
	f.args = append(f.args, &funcArg{name: name, typ: t})
}

func (f *function) Name() string        { return f.name }
func (f *function) Kind() symbol.Kind   { return symbol.Func }
func (f *function) Token() *lexer.Token { return f.token }
