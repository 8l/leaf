package types

type Basic byte

const (
	Int32 Basic = iota
	Uint32
	Int8
	Uint8
	Float64
)

func (self Basic) Size() uint32 {
	switch self {
	case Int32, Uint32:
		return 4
	case Int8, Uint8:
		return 1
	case Float64:
		return 8
	default:
		panic("bug")
	}
}

var basicNames = map[Basic]string{
	Int32:  "int",
	Uint32: "uint",
	Int8:   "char",
	Uint8:  "byte",
}

func (self Basic) String() string {
	return basicNames[self]
}
