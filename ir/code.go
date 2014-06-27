package ir

import (
	"e8vm.net/e8/inst"
	sym "e8vm.net/leaf/ir/symbol"
)

type Code struct {
	table     *sym.Table
	frameSize int16 // from SP to SP + framesize
	argsSize  int16 // from SP to SP - argssize
	insts     []Inst

	retAddr StackObj
	ret     StackObj
	args    []StackObj
}

const (
	regSP  = 29          // stack pointer
	regRet = inst.RegRet // return address
	regPC  = inst.RegPC
)

func (self *Code) Size() uint32 {
	return uint32(len(self.insts)) * 4
}

func (self *Code) Query(name string) Obj {
	s := self.table.Get(name)
	if s == nil {
		return nil
	}

	switch s.Class() {
	case sym.Func:
		// must be a top function declaration
		f := s.Obj().(*Func)
		return &Sym{f.file.pack.path, f.name}
	default:
		panic("todo")
	}
}

func (self *Code) fetchArg(size int16) StackObj {
	var ret StackObj
	self.argsSize += alignUp(size)
	ret.Offset = -self.argsSize
	ret.Len = size

	return ret
}

func alignUp(size int16) int16 {
	if size%4 != 0 {
		size = ((size >> 2) << 2) + 4
	}
	return size
}

func (self *Code) push(size int16) StackObj {
	assert(size > 0)

	var ret StackObj
	ret.Offset = self.frameSize
	ret.Len = size

	self.frameSize += alignUp(size)
	// check frameSize overflow here
	return ret
}

func (self *Code) pushReg(r uint8) StackObj {
	ret := self.push(4)
	self.swStack(r, ret)
	return ret
}

// Allocate a new stack object.
func (self *Code) Var(name string, size int16) StackObj {
	ret := self.push(size)
	if name != "_" {
		// register a variable symbol here
		panic("todo")
	}
	return ret
}

// Push an object onto the top of the stack frame,
// increases the stack frame size, and returns the new stack object.
func (self *Code) Push(o Obj) StackObj {
	assert(o != nil)

	switch o := o.(type) {
	case Imm:
		self.loadi(1, uint32(o))
		return self.pushReg(1)
	case *Sym:
		self.loadiSym(1, o)
		return self.pushReg(1)
	case StackObj:
		panic("todo")
	case HeapObj:
		panic("todo")
	default:
		panic("bug")
	}
}

// Remove temp objects from the stack frame,
// decreases the stack frame size
func (self *Code) Pop(objs ...Obj) {
	size := int16(0)
	for _, o := range objs {
		switch o := o.(type) {
		case Imm, *Sym:
			size += 4
		case StackObj:
			size += o.Len
		default:
			panic("bug")
		}
	}
	self.frameSize -= size
	assert(self.frameSize >= 0)
}

// Call a function object.
func (self *Code) Call(o Obj) {
	assert(o != nil)

	switch o := o.(type) {
	case Imm:
		// should never need to call an immediate
		// should always call a symbol
		panic("bug")
	case *Sym:
		self.addi(regSP, regSP, self.frameSize) // move to the next frame
		self.jalSym(o)                          // perform the call
		self.subi(regSP, regSP, self.frameSize) // move back to this frame
	case StackObj: // a function variable
		panic("todo")
	case HeapObj: // a function variable
		panic("todo")
	default:
		panic("bug")
	}
}

// Jump back to the calling PC position.
func (self *Code) Return() {
	self.lwStack(regPC, self.retAddr)
}
