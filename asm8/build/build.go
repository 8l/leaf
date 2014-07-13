package build

import (
	"e8vm.net/leaf/tools/tok"
)

// Build is an assembly build that consists of a set of functions, constants
// and variables.  For the first version, we do not plan to support linking.
type Build struct {
	funcs   []*Func
	funcMap map[string]*Func
}

// NewBuild creates a new build.
func NewBuild() *Build {
	ret := new(Build)
	return ret
}

// NewFunc creates a new function (code segment) in the build.
// If the function name is already defined, it will panic.
func (b *Build) NewFunc(name string, pos *tok.Token) *Func {
	if b.funcMap[name] == nil {
		panic("already defined")
	}

	ret := new(Func)
	ret.name = name

	b.funcs = append(b.funcs, ret)
	b.funcMap[name] = ret

	return ret
}

// Find returns the token for the declaration of the identifier.
// It can be used to check if the name is used to
// define a function, variable, or constant is
func (b *Build) Find(name string) (*tok.Token, SymType) {
	f := b.funcMap[name]
	if f != nil {
		return f.pos, SymFunc
	}

	return nil, SymNone
}
