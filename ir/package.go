package ir

type Package struct {
	Name string

	Deps []*Dep
	This *Dep

	Funcs []*Func
	Vars  []*Var

	depMap map[string]*Dep
}

func NewPackage(name string) *Package {
	ret := new(Package)
	ret.Name = name
	ret.This = new(Dep)
	ret.This.Name = name

	return ret
}

func (self *Package) NewVar() *Var {
	ret := new(Var)
	self.Vars = append(self.Vars, ret)

	ret.Index = len(self.Vars)
	ret.IsHeap = true

	return ret
}

func (self *Package) NewDep(name string) *Dep {
	if self.depMap[name] != nil {
		panic("bug")
	}

	ret := new(Dep)
	ret.Name = name

	self.depMap[name] = ret

	return ret
}
