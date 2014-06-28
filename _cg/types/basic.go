package types

type basic int

const (
	V basic = iota
	I32
	U32
	I8
	U8
)

func (b basic) Size() uint32 {
	switch b {
	case V:
		return 0
	case I32, U32:
		return 4
	case I8, U8:
		return 1
	}

	panic("bug")
}
