package build

import (
	"bytes"

	"e8vm.net/leaf/tools/tok"
)

// Func is a continuous memory segment that saves executable instructions.
type Func struct {
	pos    *tok.Token
	name   string
	lines  []*Line
	labels map[string]*label
}

func newFunc(name string, pos *tok.Token) *Func {
	ret := new(Func)
	ret.name = name
	ret.pos = pos
	ret.labels = make(map[string]*label)

	return ret
}

// NewLine appends a new empty line in the function to fill.
func (f *Func) NewLine() *Line {
	ret := new(Line)
	f.lines = append(f.lines, ret)
	return ret
}

// FindLabel returns the token of a previously marked label.
// It retunrs nil if the label is not marked yet.
func (f *Func) FindLabel(label string) *tok.Token {
	get := f.labels[label]
	if get == nil {
		return nil
	}
	return get.pos
}

// LocateLabel returns the offset location of the marked label
// It returns -1 if the label is not marked yet.
func (f *Func) LocateLabel(label string) int {
	lab := f.labels[label]
	if lab == nil {
		return -1
	}

	return lab.loc
}

// MarkLabel marks a label at the current writing position.
func (f *Func) MarkLabel(pos *tok.Token) {
	name := pos.Lit
	assert(f.labels[name] == nil)
	lab := &label{pos, len(f.lines)}
	f.labels[name] = lab
}

// emit writes out the code as it is.
// to have meaningful code, the build has to be generated first.
func (f *Func) emit(buf *bytes.Buffer) {
	b := make([]byte, 4)

	for _, line := range f.lines {
		i := uint32(line.Inst)

		b[0] = uint8(i)
		b[1] = uint8(i >> 8)
		b[2] = uint8(i >> 16)
		b[3] = uint8(i >> 24)

		_, e := buf.Write(b)
		if e != nil {
			panic("bug")
		}
	}
}
