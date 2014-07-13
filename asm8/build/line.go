package build

import (
	"e8vm.net/e8/inst"
)

// Line is an instruction where the lable might be not binded. Each
// instruction always take exactly one uint32.
type Line struct {
	i     inst.Inst
	label string // how the label is applied depends on the inst op
}

func newLine(i inst.Inst) *Line {
	ret := new(Line)
	ret.i = i
	return ret
}

func newLineSym(i inst.Inst, lab string) *Line {
	ret := newLine(i)
	ret.label = lab
	return ret
}
