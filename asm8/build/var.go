package build

import (
	"bytes"

	"e8vm.net/leaf/tools/tok"
)

// Var describes a data segment
type Var struct {
	*bytes.Buffer

	pos   *tok.Token
	name  string
	align uint32
	pad   uint32
}

func newVar(name string, pos *tok.Token) *Var {
	ret := new(Var)
	ret.name = name
	ret.pos = pos

	ret.Buffer = new(bytes.Buffer)

	return ret
}

// Align sets the alignment of the start of the data segment.
func (v *Var) Align(a uint32) {
	if a != 0 && a != 4 && a != 8 {
		panic("invalid alignment")
	}

	v.align = a
}

// AlignStart calculates the start address of the segment
// by advancing the address.
func (v *Var) AlignStart(start uint32) uint32 {
	if v.align == 0 {
		v.pad = 0
		return start
	}

	res := start % v.align
	if res == 0 {
		v.pad = 0
		return start
	}
	v.pad = v.align - res
	return start + v.align - res
}

func (v *Var) emit(buf *bytes.Buffer) {
	if v.pad > 0 {
		buf.Write(make([]byte, v.pad))
	}
	buf.Write(v.Buffer.Bytes()) // copy out
}
