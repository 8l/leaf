package ir

type Build struct {
	packs    []*Package
	packMap  map[string]*Package
	typeList *typeList
}

func NewBuild() *Build {
	ret := new(Build)
	ret.packMap = make(map[string]*Package)
	ret.typeList = newTypeList()

	ret.addBuiltIn()

	return ret
}

func (self *Build) addBuiltIn() {
	p := newPackage("<builtin>")
	self.addPackage(p)

	void := p.TypeRef(nil)
	assert(void.importId == 0 && void.typeId == 0)

	u32 := p.TypeRef(Uint32)
	i32 := p.TypeRef(Int32)
	u8 := p.TypeRef(Uint8)
	i8 := p.TypeRef(Int8)

	ptr := p.TypeRef(p.NewPointerType(void))

	p.DeclType("int32", i32)
	p.DeclType("uint32", u32)
	p.DeclType("int", i32)
	p.DeclType("uint", u32)
	p.DeclType("int8", i8)
	p.DeclType("uint8", u8)
	p.DeclType("char", i8)
	p.DeclType("byte", u8)
	p.DeclType("ptr", ptr)

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

func (self *Func) Define() *Code {
	ret := new(Code)
	self.code = ret
	return ret
}
