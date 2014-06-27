package ir

type Obj interface{}
type Imm uint32
type Sym struct{ Pack, Name string }
type StackObj struct {
	Offset int16
	Len    int16
}
type HeapObj struct{ Addr, Len uint32 }
