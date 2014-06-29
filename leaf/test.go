package main

import (
	"os"

	"e8vm.net/leaf/ir"
	"e8vm.net/leaf/ir/types"
)

func mainTest(_ []string) {
	b := ir.NewBuild()

	p := b.NewPackage("e8")
	file := p.NewFile("e8.leaf")
	fn, _ := file.DeclNewFunc("main", types.NewFunc(nil))

	c := fn.Define()
	i := c.Push(ir.Const(int64('E'), types.Uint8))
	pch := c.QueryObj("printChar")
	c.Call(pch)
	c.Pop(i)

	i = c.Push(ir.Const(int64('8'), types.Uint8))
	c.Call(pch)
	c.Pop(i)

	i = c.Push(ir.Const(int64('\n'), types.Uint8))
	c.Call(pch)
	c.Pop(i)

	c.Return()

	// p.Save()
	b.Print()

	fout, e := os.Create("out.e8")
	if e != nil {
		panic(e)
	}
	b.Build("e8", fout)
	e = fout.Close()
	if e != nil {
		panic(e)
	}
}
