package build

import (
	"e8vm.net/e8/inst"
)

// line is an instruction where the label might be not binded. Each
// instruction always take exactly one uint32.
type line struct {
	i     inst.Inst
	label string // how the label is applied depends on the inst op
}

func newLine(i inst.Inst) *line {
	ret := new(line)
	ret.i = i
	return ret
}

func newLineSym(i inst.Inst, lab string) *line {
	ret := newLine(i)
	ret.label = lab
	return ret
}
