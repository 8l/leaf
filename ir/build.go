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

func (self *Build) LoadPackage(path string) *Package {
	panic("todo")
}
