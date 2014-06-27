package ir

// Package is only a collection of stuff
// It does not provide name to object mapping.
// An object might have a name, but it is only for debugging
type Package struct {
	Name string

	This *Dep
	Deps []*Dep

	Funcs []*Func
	Vars  []*Var
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
	ret := new(Dep)
	ret.Name = name

	self.Deps = append(self.Deps, ret)
	return ret
}
