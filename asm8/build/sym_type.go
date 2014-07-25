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

func (t SymType) String() string {
	switch t {
	case SymNone:
		return "<none>"
	case SymFunc:
		return "function"
	case SymVar:
		return "variable"
	case SymConst:
		return "const"
	}

	return "<unknown>"
}
