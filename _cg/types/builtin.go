package types

var (
	Void   = newBuiltIn("void", V)
	Int    = newBuiltIn("int", I32)
	Uint   = newBuiltIn("uint", U32)
	Int8   = newBuiltIn("int8", I8)
	Uint8  = newBuiltIn("uint8", U8)
	Int32  = newBuiltIn("int32", I32)
	Uint32 = newBuiltIn("uint32", U32)
	Byte   = newBuiltIn("byte", U8)
	Char   = newBuiltIn("char", I8)

	Ptr    = newBuiltIn("ptr", NewPointer(V))
	String = newBuiltIn("string", NewPointer(I8)) // might change
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
