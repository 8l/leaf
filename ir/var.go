package ir

type Var struct {
	Index   int
	Offset  uint32
	Size    uint32
	IsConst bool
	IsHeap  bool

	Value uint32 // for constants

	name string
}
