package types

var (
	Void   = newBuiltIn("void", _void)
	Int    = newBuiltIn("int", _int32)
	Uint   = newBuiltIn("uint", _uint32)
	Int8   = newBuiltIn("int8", _int8)
	Uint8  = newBuiltIn("uint8", _uint8)
	Int32  = newBuiltIn("int32", _int32)
	Uint32 = newBuiltIn("uint32", _uint32)
	Byte   = newBuiltIn("byte", _uint8)
	Char   = newBuiltIn("char", _int8)

	Ptr    = newBuiltIn("ptr", NewPointer(_void))
	String = newBuiltIn("string", NewPointer(_int8)) // might change
)

var BuiltIns = []*Named{
	Void,
	Int, Uint,
	Int8, Uint8,
	Int32, Uint32,
	Byte, Char,
	Ptr,
	String,
}
