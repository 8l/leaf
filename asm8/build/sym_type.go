package build

// SymType reflects the type of a symbol
type SymType int

// A list of all the symbol types.
const (
	SymNone SymType = iota
	SymFunc
	SymConst
	SymVar
)
