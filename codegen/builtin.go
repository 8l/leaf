package codegen

var (
	tpVoid   = defType("void", _void)
	tpUint   = defType("uint", _uint32)
	tpInt    = defType("int", _int32)
	tpUint8  = defType("uint8", _uint8)
	tpInt8   = defType("int8", _int8)
	tpInt32  = defType("int32", _int32)
	tpUint32 = defType("uint32", _uint32)
	tpByte   = defType("byte", _uint8)
	tpChar   = defType("char", _int8)

	tpPtr    = defType("ptr", ptr(_void))
	tpString = defType("string", ptr(_int8))

	fnPrintInt *function
	fnPrintStr *function
)

func makeBuiltIn() *symMap {
	ret := newSymMap()

	ret.Add(tpUint, tpInt)
	ret.Add(tpUint8, tpInt8)
	ret.Add(tpUint32, tpInt32)
	ret.Add(tpByte, tpChar)

	ret.Add(fnPrintInt) // printInt
	ret.Add(fnPrintStr) // printStr

	return ret
}

var builtIn *symMap = makeBuiltIn()

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
