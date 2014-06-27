package ir

import (
	"e8vm.net/e8/inst"
	sym "e8vm.net/leaf/ir/symbol"
)

type Code struct {
	table     *sym.Table
	frameSize int16
	insts     []Inst
}

/*
$29: stack top pointer, always increase
$30: stack base pointer
$31: pc
*/

const (
	regSP = 30 // stack pointer
)

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

func (self *Code) push(size int16) StackObj {
	var ret StackObj
	ret.Offset = self.frameSize
	ret.Len = size

	self.frameSize += size
	// check frameSize overflow here
	return ret
}

func (self *Code) pushReg(r uint8) StackObj {
	ret := self.push(4)
	inst.Iinst(inst.OpSw, regSP, r, uint16(ret.Offset))
	return ret
}

func (self *Code) Push(o Obj) Obj {
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
}

func (self *Code) Call(o Obj) {
	assert(o != nil)

	switch o.(type) {
	case Imm:
		panic("bug") // remove this when you really need to call an imm
	case *Sym:
		panic("todo")
	case StackObj: // a function variable
		panic("todo")
	case HeapObj: // a function variable
		panic("todo")
	default:
		panic("bug")
	}
}

func (self *Code) Return(f Obj)     {}
func (self *Code) Imm(i uint32) Obj { return Imm(i) }

/*
A code segment is a block of IR.
This might be changed in the future, but for now, we have
several types of objects

two types of immediats
- constant, a uint32 immediate constant
- symbol, a uint32 immediate constant that will be filled later on linking

two types of memory objects
- heapObj, a memory block that stays on the heap, a data segment
- stackObj, a memory block that stays on the stack

IR does not care about type.

IR instructions:

// first will declare the stack objects
// though currently there is no stack objects needed

push <obj> // hard copy an object to the calling stack
call <obj> // increase the stack, jump to the address at the object value
pop // restore the call stack
ret // return function call
assign <obj1> <obj2> // copy the value of the obj2 to the location of obj1

*/
