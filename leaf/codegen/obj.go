package codegen

import (
	"e8vm.net/leaf/leaf/ir"
	"e8vm.net/leaf/leaf/ir/types"
)

type obj struct {
	o ir.Obj
	t types.Type
}

var voidObj *obj = new(obj)
