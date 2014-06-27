package ir

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
	types    []int // map to type id in the package type list
	symbols  []Symbol
	dataSegs []*Data
	codeSegs []*Code
}

func (self *Package) NewFile(name string) *File {
	ret := new(File)
	ret.pack = self
	ret.name = name
	return ret
}

func (self *Package) Save() {
	panic("todo")
}

// Based on a type object finds its equivalent type ref.
// If the type object is a new unique one, then add it
// into the type object library.
func (self *Package) TypeRef(t Type) TypeRef {
	panic("todo")
}

// Add a new function type
func (self *Package) NewFuncType() *FuncType {
	ret := new(FuncType)
	ret.pack = self
	return ret
}

func (self *Package) NewPointerType(t TypeRef) *PointerType {
	ret := new(PointerType)
	ret.pack = self
	ret.t = t
	return ret
}

func (self *Package) DeclType(name string, tr TypeRef) Symbol {
	panic("todo")
}
