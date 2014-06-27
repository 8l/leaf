package main

import (
	"e8vm.net/leaf/ir"
	"e8vm.net/leaf/ir/types"
)

func mainTest(_ []string) {
	b := ir.NewBuild()
	p := b.NewPackage("p")
	file := p.NewFile("a.leaf")
	fn := file.NewFunc("main", types.NewFunc(nil))
	file.DeclFunc(fn)

	seg := fn.Define()
	seg.Push(seg.Imm(42))
	seg.Call(seg.Query("printInt"))

	p.Save()
}
