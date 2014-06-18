package codegen

var (
	VoidType = &BasicType{"void", Void}

	PtrType    = &BasicType{"ptr", Ptr}
	UintType   = &BasicType{"uint", Uint32}
	IntType    = &BasicType{"int", Int32}
	Uint8Type  = &BasicType{"uint8", Uint8}
	Int8Type   = &BasicType{"int8", Int8}
	Int32Type  = &BasicType{"int32", Int32}
	Uint32Type = &BasicType{"uint32", Uint32}
	ByteType   = &BasicType{"byte", Uint8}
	CharType   = &BasicType{"char", Int8}

	fnPrintInt *Func
	// fnPrintStr *Func
)

func makeBuiltIn() *SymTable {
	ret := newSymTable()

	ret.Add(PtrType)
	ret.Add(UintType, IntType)
	ret.Add(Uint8Type, Int8Type)
	ret.Add(Uint32Type, Int32Type)
	ret.Add(ByteType, CharType)

	ret.Add(fnPrintInt)

	return ret
}

func init() {
	fnPrintInt = func() *Func {
		f := NewFunc("printInt")
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
