package header

// Stores a package header
type Package struct {
	Imports []string // list of imports, all in absolute import path, 0 is always builtin
	Symbols []Symbol // all symbol

	Numbers []*Number // constant numbers
	Strings []*String // constant strings

	Structs []*Struct // structures
	// Funcs   []*Func   // function signatures
	// Vars    []*Var    // variables
}

type Symbol interface {
}

type SymRef struct {
	Package int // package index, 0 for builtin, -1 for this package
	Index   int // symbol index
}

type Number struct {
	Value int64
}

type String struct {
	Value string
}

type Struct struct {
	Fields []*Field
}

type Field struct {
	Name string
	Type Type
}

type Type interface{}

type Pointer struct {
	Type Type
}

type Slice struct {
	Type Type
}

type Array struct {
	Type Type
	Size int
}

type Basic struct {
	Name string
	int
}
