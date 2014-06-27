package ir

func assert(cond bool) {
	if !cond {
		panic("bug")
	}
}
