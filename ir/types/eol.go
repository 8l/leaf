package types

// end of list
type eol int

func (self eol) Size() uint32 { panic("bug") }

func IsEOL(t Type) bool {
	_, ret := t.(eol)
	return ret
}
