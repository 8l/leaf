package codegen

type basicType int

// basic
const (
	Void basicType = iota
	Int32
	Uint32
	Int8
	Uint8
)

var _ Type = Void

func (b basicType) Size() uint32 {
	switch b {
	case Void:
		return 0
	case Int32:
		return 4
	case Uint32:
		return 4
	case Int8:
		return 1
	case Uint8:
		return 1
	default:
		panic("bug")
	}
}

func (b basicType) String() string {
	switch b {
	case Void:
		return "<void>"
	case Int32:
		return "<int32>"
	case Uint32:
		return "<uint32>"
	case Int8:
		return "<int8>"
	case Uint8:
		return "<uint8>"
	default:
		panic("bug")
	}
}

func isBasic(t Type) bool {
	_, ret := t.(basicType)
	return ret
}
