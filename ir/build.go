package ir

type Build struct {
	builtIn *Package

	packs   []*Package
	packMap map[string]*Package
}

func NewBuild() *Build {
	ret := new(Build)
	ret.packMap = make(map[string]*Package)

	ret.builtIn = makeBuiltIn()
	ret.addPackage(ret.builtIn)

	return ret
}

func (self *Build) addPackage(p *Package) {
	if self.packMap[p.path] != nil {
		panic("package with the same path already exists")
	}

	p.build = self
	self.packs = append(self.packs, p)
	self.packMap[p.path] = p
}

func (self *Build) NewPackage(path string) *Package {
	ret := newPackage(path)
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

const (
	segSize = (1 << 32) / 4
)

const (
	ioStart uint32 = segSize * iota
	codeStart
	heapStart
	stackStart
)

func (self *Build) Build(p string) {
	c := new(Code)
	c.loadi(regSP, stackStart) // init the stack pointer
	c.jalSym(&Sym{p, "main"})  // jump to the main function
	c.sb(0, 0, ioHalt)         // halt the VM

	panic(c)
}
