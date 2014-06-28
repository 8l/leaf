package ir

import (
	"e8vm.net/leaf/ir/types"
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

func ConstNum(i int64) Imm {
	return Imm{i, types.ConstNum}
}

func ConstInt(i int64, t types.Basic) Imm {
	assert(t != types.Float64)
	assert(t != types.ConstNum)

	return Imm{i, t}
}
