package codegen

import (
	"e8vm.net/leaf/codegen/symbol"
	"e8vm.net/leaf/codegen/types"
)

var (
	fnPrintInt *function
	fnPrintStr *function
)

func init() {
	fnPrintInt = func() *function {
		f := newBuiltInFunc("printInt")
		f.AddArg(types.Int)
		return f
	}()

	fnPrintStr = func() *function {
		f := newBuiltInFunc("printStr")
		f.AddArg(types.String)
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
