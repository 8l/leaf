package ir

import (
	"e8vm.net/leaf/leaf/ir/types"
)

type Obj interface{}

type Imm struct {
	Value int64
	Type  types.Basic
}

type Sym struct {
	Pack, Name string
}

type StackObj struct {
	Offset, Len int16
}

type HeapObj struct {
	Addr, Len uint32
}

func Const(i int64, t types.Basic) Imm {
	assert(t != types.Float64)

	return Imm{i, t}
}
