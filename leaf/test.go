package main

import (
	"e8vm.net/leaf/ir"
	"e8vm.net/leaf/ir/types"
)

func mainTest(_ []string) {
	b := ir.NewBuild()

	p := b.NewPackage("e8")
	file := p.NewFile("e8.leaf")
	fn := file.NewFunc("main", types.NewFunc(nil))
	file.DeclFunc(fn)

	c := fn.Define()
	i := c.Push(ir.Imm(uint8('E')))
	c.Call(c.Query("printChar"))
	c.Pop(i)

	i = c.Push(ir.Imm(uint8('8')))
	c.Call(c.Query("printChar"))
	c.Pop(i)

	i = c.Push(ir.Imm(uint8('\n')))
	c.Call(c.Query("printChar"))
	c.Pop(i)

	c.Return()

	// p.Save()
	b.Print()
	b.Build("e8")
}
