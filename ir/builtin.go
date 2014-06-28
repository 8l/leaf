package ir

import (
	"e8vm.net/e8/vm"
	"e8vm.net/leaf/ir/types"
)

func makeBuiltIn() *Package {
	p := newPackage("builtin")
	f := p.NewFile("builtin.leaf")

	f.DeclType("int32", types.Int32)
	f.DeclType("uint32", types.Uint32)
	f.DeclType("int", types.Int32)
	f.DeclType("uint", types.Uint32)
	f.DeclType("int8", types.Int8)
	f.DeclType("uint8", types.Uint8)
	f.DeclType("char", types.Int8)
	f.DeclType("byte", types.Uint8)
	f.DeclType("ptr", types.NewPointer(nil))

	// f.DeclFunc(f.NewFunc("printInt", types.NewFunc(nil, types.Int32)))

	def := func(name string, t *types.Func, fn func(*Func)) {
		ret := f.NewFunc(name, t)
		f.DeclFunc(ret)
		fn(ret)
	}

	def("putc", types.NewFunc(nil, types.Int8), _printChar)
	def("printChar", types.NewFunc(nil, types.Int8), _printChar)
	def("printInt", types.NewFunc(nil, types.Int32), _printInt)

	return p
}

func _printChar(f *Func) {
	c := f.Define()

	c.lbuStack(1, c.args[0]) // load the parameter to $1
	c.lbu(2, 0, vm.Stdout)   // check if output is ready using $2
	c.bne(2, 0, -2)          // keep pulling if not ready
	c.sb(1, 0, vm.Stdout)    // write the byte out

	c.Return()
}

func _printInt(f *Func) {
	c := f.Define()
	// TODO
	c.Return()
}

func makeEntry(pack string) *Code {
	c := new(Code)
	c.loadi(regSP, stackStart)   // init the stack pointer
	c.jalSym(&Sym{pack, "main"}) // jump to the main function
	c.sb(0, 0, vm.Halt)          // halt the VM
	return c
}
