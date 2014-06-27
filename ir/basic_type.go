package ir

type BasicType int

const (
	Int32 BasicType = iota
	Uint32
	Int8
	Uint8
	Float64
)

func (self BasicType) Size() uint32 {
	switch self {
	case Uint32, Int32:
		return 4
	case Int8, Uint8:
		return 1
	case Float64:
		return 8
	}
	panic("bug")
}
