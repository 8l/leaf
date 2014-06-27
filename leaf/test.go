package main

import (
	"e8vm.net/leaf/ir"
	"e8vm.net/leaf/ir/types"
)

func mainTest(_ []string) {
	b := ir.NewBuild()
	p := b.NewPackage("p")
	file := p.NewFile("a.leaf")
	typ := new(types.Func)
	fn := file.DeclareFunc("main", typ)

	seg := fn.Define()
	v1 := seg.Number(42)
	seg.Push(v1)
	v2 := seg.Query("printInt")
	seg.Call(v2)

	p.Save()
}
