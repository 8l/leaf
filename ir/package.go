package ir

func makeBuiltIn() *Package {
	ret := newPackage("<builtin>")
	panic("todo")
	// TODO: add builtin types and funtions here
	return ret
}

func newPackage(path string) *Package {
	ret := new(Package)
	ret.path = path
	return ret
}

type Package struct {
	path  string // absolute path of the package
	pid   int
	build *Build

	imports  []int // packageIds
	types    []Type
	symbols  []Symbol
	dataSegs []*Data
	codeSegs []*Code
}

func (self *Package) NewFile(name string) *File {
	ret := new(File)
	ret.pack = self
	return ret
}

func (self *Package) Save() {
	panic("todo")
}

// Based on a type object finds its equivalent type ref.
// If the type object is a new unique one, then add it
// into the type object library.
func (self *Package) typeRef(t Type) TypeRef {
	panic("todo")
}

// Add a new function type
func (self *Package) NewFuncType() *FuncType {
	ret := new(FuncType)
	ret.pack = self
	return ret
}

func (self *Package) TypeRef(t Type) TypeRef {
	panic("todo")
}
