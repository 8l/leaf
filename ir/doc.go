package ir

/*
	// on linking stage, we don't need the types
	// so on linking stage, function printInt is only the
	// starting address of a code block
	// also for this extern code block, we don't need to know its actual size

	// external symbols
	package <builtin>:
		func printInt

	func main:
		// no arg
		// no ret
		<0> = 42
		<1> = link <builtin>.printInt
		push <0>
		call <1>

	and from a building perspective

	f := NewFunc()
	_0 := f.Const(42)
	_1 := f.Link(printInt)
	f.Push(_0)
	f.Call(_1)

*/
