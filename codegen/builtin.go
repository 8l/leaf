package codegen

import (
	"e8vm.net/leaf/symbol"
	"e8vm.net/leaf/types"
)

var (
	fnPrintInt *function
	fnPrintStr *function
)

func init() {
	fnPrintInt = func() *function {
		f := newBuiltInFunc("printInt")
		f.addArg(types.Int)
		return f
	}()

	fnPrintStr = func() *function {
		f := newBuiltInFunc("printStr")
		f.addArg(types.String)
		return f
	}()
}

func makeBuiltIn() *symbol.Scope {
	ret := symbol.NewScope()

	for _, s := range types.BuiltIns {
		ret.Register(s)
	}

	ret.Register(fnPrintInt)
	ret.Register(fnPrintStr)

	return ret
}

var builtIn *symbol.Scope = makeBuiltIn()
