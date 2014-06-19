package codegen

var (
	VoidType   = &bacicType{"void", Void}
	UintType   = &bacicType{"uint", Uint32}
	IntType    = &bacicType{"int", Int32}
	Uint8Type  = &bacicType{"uint8", Uint8}
	Int8Type   = &bacicType{"int8", Int8}
	Int32Type  = &bacicType{"int32", Int32}
	Uint32Type = &bacicType{"uint32", Uint32}
	ByteType   = &bacicType{"byte", Uint8}
	CharType   = &bacicType{"char", Int8}

	StrType = &NamedType{"string", &PtrType{CharType}}

	fnPrintInt *function
	// fnPrintStr *Func
)

func makeBuiltIn() *symMap {
	ret := newSymMap()

	ret.Add(UintType, IntType)
	ret.Add(Uint8Type, Int8Type)
	ret.Add(Uint32Type, Int32Type)
	ret.Add(ByteType, CharType)

	ret.Add(fnPrintInt)

	return ret
}

func init() {
	fnPrintInt = func() *function {
		f := newFunc("printInt")
		f.Ret = VoidType
		f.AddArg(IntType)
		return f
	}()

	/*
		fnPrintStr = func() *Func {
			f := NewFunc("printStr")
			f.Ret = VoidType
			f.AddArg(StrType)
			return f
		}
	*/
}
