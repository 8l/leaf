package codegen

import (
	"e8vm.net/leaf/symbol"
)

func builtInType(name string, t typ) *namedType {
	return &namedType{name, nil, t}
}

var (
	tpVoid   = builtInType("void", _void)
	tpUint   = builtInType("uint", _uint32)
	tpInt    = builtInType("int", _int32)
	tpUint8  = builtInType("uint8", _uint8)
	tpInt8   = builtInType("int8", _int8)
	tpInt32  = builtInType("int32", _int32)
	tpUint32 = builtInType("uint32", _uint32)
	tpByte   = builtInType("byte", _uint8)
	tpChar   = builtInType("char", _int8)

	tpPtr    = builtInType("ptr", ptr(_void))
	tpString = builtInType("string", ptr(_int8))

	fnPrintInt *function
	fnPrintStr *function
)

func makeBuiltIn() *symbol.Scope {
	syms := []symbol.Symbol{
		tpInt, tpUint,
		tpInt8, tpUint8,
		tpInt32, tpUint32,
		tpByte, tpChar,
		fnPrintInt,
		fnPrintStr,
	}

	ret := symbol.NewScope()
	for _, s := range syms {
		ret.Register(s)
	}

	return ret
}

var builtIn *symbol.Scope = makeBuiltIn()

func init() {
	fnPrintInt = func() *function {
		f := newFunc("printInt")
		f.addArg(tpInt)
		return f
	}()

	fnPrintStr = func() *function {
		f := newFunc("printStr")
		f.addArg(tpString)
		return f
	}()
}
