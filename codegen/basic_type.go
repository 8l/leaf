package codegen

type basicType int

// internal basic types
// you will never be able to reference these types
// but they exist in the abstract world
const (
	_void basicType = iota
	_int32
	_uint32
	_int8
	_uint8
)

var _ typ = _void

func (b basicType) Size() uint32 {
	switch b {
	case _void:
		return 0
	case _int32:
		return 4
	case _uint32:
		return 4
	case _int8:
		return 1
	case _uint8:
		return 1
	default:
		panic("bug")
	}
}

func (b basicType) String() string {
	switch b {
	case _void:
		return "<void>"
	case _int32:
		return "<int32>"
	case _uint32:
		return "<uint32>"
	case _int8:
		return "<int8>"
	case _uint8:
		return "<uint8>"
	default:
		panic("bug")
	}
}

func isBasic(t typ) bool {
	_, ret := t.(basicType)
	return ret
}
