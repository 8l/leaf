package types

type basic int

const (
	_void basic = iota
	_int32
	_uint32
	_int8
	_uint8
)

func (b basic) Size() uint32 {
	switch b {
	case _void:
		return 0
	case _int32, _uint32:
		return 4
	case _int8, _uint8:
		return 1
	}

	panic("bug")
}
