package ir

type Obj interface{}
type Imm uint32
type Sym struct{ Pack, Name string }
type StackObj struct{ Offset, Len uint32 }
type HeapObj struct{ Addr, Len uint32 }
