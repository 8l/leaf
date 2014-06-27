package ir

type Build struct {
	packs   []*Package
	packMap map[string]*Package
}

func NewBuild() *Build {
	ret := new(Build)
	builtin := makeBuiltIn()

	ret.addPackage(builtin)

	return ret
}

func (self *Build) addPackage(p *Package) {
	if self.packMap[p.path] != nil {
		panic("package with the same path already exists")
	}

	p.pid = len(self.packs)
	p.build = self
	self.packs = append(self.packs, p)
	self.packMap[p.path] = p
}

func (self *Build) NewPackage(path string) *Package {
	ret := new(Package)
	ret.path = path
	self.addPackage(ret)

	return ret
}

func (self *Build) ImportPackage(path string) *Package {
	if self.packMap[path] == nil {
		self.importPackage(path)
	}
	return self.packMap[path]
}

// import the header (consts, types and symbols) only
func (self *Build) importPackage(path string) *Package {
	panic("todo")
}

func (self *Build) LoadPackage(path string) *Package {
	panic("todo")
}

func makeBuiltIn() *Package {
	ret := newPackage("<builtin>")
	// add builtin types and funtions here
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
	dataSegs []*DataSeg
	codeSegs []*CodeSeg
}

func (self *Package) NewFile(name string) *File {
	ret := new(File)
	ret.pack = self
	return ret
}

func (self *Package) Save() {
}

func (self *Package) typeRef(t Type) TypeRef {
	panic("todo")
}

func (self *Package) NewFuncType() TypeRef {
	panic("todo")
}

type File struct {
	pack    *Package
	imports map[string]int
}

func (self *File) DeclareFunc(name string, t TypeRef) *Func {
	ret := new(Func)
	ret.name = name
	ret.t = t
	ret.file = self
	// TODO add symbol into package
	return ret
}

func (self *Func) Define() *CodeSeg {
	ret := new(CodeSeg)
	self.code = ret
	return ret
}

type Func struct {
	name string
	t    TypeRef
	file *File
	code *CodeSeg
}

type Type interface {
	Size() uint32
}

type TypeRef struct {
	importId int
	typeid   int
}

type Symbol interface {
	Name() string
}

type DataSeg struct {
	size uint32
}

type CodeSeg struct {
}

// A symbol
type Var struct {
	name string
}

type NumConst struct {
	name string
	t    Type
	v    *Num
}

type Num struct{}

type Ref struct{}

func (self *CodeSeg) Query(name string) *Ref {
	panic("todo")
}

func (self *CodeSeg) Push(v *Ref) {
}

func (self *CodeSeg) Call(f *Ref) *Ref {
	panic("todo")
}

func (self *CodeSeg) Return(f *Ref) {
}

func (self *CodeSeg) Number(i int64) *Ref {
	panic("todo")
}
