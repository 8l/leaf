package build

import (
	"e8vm.net/e8/inst"
	"e8vm.net/leaf/tools/tok"
)

// Func is a continuous memory segment that saves executable instructions.
type Func struct {
	pos    *tok.Token
	name   string
	lines  []*Line
	labels map[string]int
}

// Line is an instruction where the lable might be not binded. Each
// instruction always take exactly one uint32.
type Line struct {
	i     inst.Inst
	label string // how the label is applied depends on the inst op
}
