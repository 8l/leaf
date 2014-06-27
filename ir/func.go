package ir

import (
	"e8vm.net/leaf/ir/types"
)

type Func struct {
	name string
	t    types.Type
	file *File
	code *Code
}
