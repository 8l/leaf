package build

func assert(cond bool) {
	if !cond {
		panic("bug")
	}
}
