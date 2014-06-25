package ir

type Var struct {
	Index   int
	Offset  uint32
	Size    uint32
	IsConst bool

	Value uint32      // for constants
	Link  interface{} // for linked symbols

	name string
}
