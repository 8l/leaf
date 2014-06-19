// Pacakge header defines the exported symbol table of a package
package header

// Stores a package header
type Package struct {
	Imports []string
	Symbols []*Symbol
}

type Symbol struct {
	Name    string
	Content interface{}
}

// defines function signature
type Func struct {
	Args []*FuncArg
	Ret  *FuncArg
}

type FuncArg struct {
	Type   Type   // how to interpret
	Size   uint32 // size on stack
	Offset uint32 // relative stack position
}
