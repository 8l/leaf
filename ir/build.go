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

func (self *Func) Define() *Code {
	ret := new(Code)
	self.code = ret
	return ret
}
