package codegen

var (
	voidType   = &namedType{"void", Void}
	uintType   = &namedType{"uint", Uint32}
	intType    = &namedType{"int", Int32}
	uint8Type  = &namedType{"uint8", Uint8}
	int8Type   = &namedType{"int8", Int8}
	int32Type  = &namedType{"int32", Int32}
	uint32Type = &namedType{"uint32", Uint32}
	byteType   = &namedType{"byte", Uint8}
	charType   = &namedType{"char", Int8}

	stringType = &namedType{"string", &ptrType{Int8}}

	fnPrintInt *function
	fnPrintStr *function
)

func makeBuiltIn() *symMap {
	ret := newSymMap()

	ret.Add(uintType, intType)
	ret.Add(uint8Type, int8Type)
	ret.Add(uint32Type, int32Type)
	ret.Add(byteType, charType)

	ret.Add(fnPrintInt)

	return ret
}

func init() {
	fnPrintInt = func() *function {
		f := newFunc("printInt")
		f.Ret = voidType
		f.AddArg(intType)
		return f
	}()

	fnPrintStr = func() *function {
		f := newFunc("printStr")
		f.Ret = voidType
		f.AddArg(stringType)
		return f
	}()
}
