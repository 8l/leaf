package build

import (
	"e8vm.net/e8/inst"
)

// Line is an instruction where the label might be not binded. Each
// instruction always take exactly one uint32.
type Line struct {
	inst.Inst
	Label string // how the label is applied depends on the inst op
}
