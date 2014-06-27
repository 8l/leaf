package ir

func example() {
	b := NewBuild()

	p := b.NewPackage("p")
	file := p.NewFile("a.leaf")
	t := p.NewFuncType()

	fn := file.DeclareFunc("main", t)

	seg := fn.Define()
	v1 := seg.Number(42)
	seg.Push(v1)
	fprint := seg.Query("printInt")
	seg.Call(fprint)
	seg.Return(nil)

	p.Save()
}
