package build

// Build is an assembly build that consists of a set of functions, constants
// and variables.  For the first version, we do not plan to support linking.
type Build struct {
}

// NewBuild creates a new build.
func NewBuild() *Build {
	ret := new(Build)
	return ret
}

func (b *Build) NewFunc() *Func {
	ret := new(Func)
	return ret
}
