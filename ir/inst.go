package ir

import (
	"e8vm.net/e8/inst"
)

type Inst struct {
	inst inst.Inst
	sym  *Sym
}

func (self *Code) inst(i inst.Inst) {
	self.instSym(i, nil)
}

func (self *Code) iinst(op uint8, s, t uint8, i int16) {
	self.inst(inst.Iinst(op, s, t, uint16(i)))
}

func (self *Code) rinst(fn uint8, s, t, d uint8) {
	self.inst(inst.Rinst(s, t, d, fn))
}

func (self *Code) jinst(op uint8, s *Sym) {
	self.instSym(inst.Jinst(op, 0), s)
}

func (self *Code) instSym(i inst.Inst, s *Sym) {
	self.insts = append(self.insts, Inst{i, s})
}

func (self *Code) add(d, s, t uint8) {
	self.rinst(inst.FnAdd, s, t, d)
}

func (self *Code) mov(d, s uint8) {
	self.add(d, s, 0)
}

func (self *Code) addi(t, s uint8, i int16) {
	self.iinst(inst.OpAddi, s, t, i)
}

func (self *Code) subi(t, s uint8, i int16) {
	self.addi(t, s, -i)
}

func (self *Code) addiSym(t, s uint8, sym *Sym) {
	self.instSym(inst.Iinst(inst.OpAddi, s, t, 0), sym)
}

func (self *Code) lui(t uint8, i uint16) {
	self.iinst(inst.OpLui, 0, t, int16(i))
}

func (self *Code) luiSym(t uint8, sym *Sym) {
	self.instSym(inst.Iinst(inst.OpLui, 0, t, 0), sym)
}

func (self *Code) bne(s, t uint8, i int16) {
	self.iinst(inst.OpBne, s, t, i)
}

func (self *Code) beq(s, t uint8, i int16) {
	self.iinst(inst.OpBeq, s, t, i)
}

func (self *Code) sb(t, s uint8, i int16) {
	self.iinst(inst.OpSb, s, t, i)
}

func (self *Code) loadi(reg uint8, i uint32) {
	up := uint16(i >> 16)
	if up != 0 {
		self.lui(reg, uint16(i>>16))
	}
	self.addi(reg, 0, int16(uint16(i)))
}

func (self *Code) loadiSym(reg uint8, sym *Sym) {
	self.luiSym(reg, sym)
	self.addiSym(reg, 0, sym)
}

func (self *Code) jSym(sym *Sym) {
	self.jinst(inst.OpJ, sym)
}

func (self *Code) jalSym(sym *Sym) {
	self.jinst(inst.OpJal, sym)
}

func (self *Code) jr(reg uint8) {
	self.mov(regPC, reg)
}

func (self *Code) lw(t, s uint8, im int16) {
	self.iinst(inst.OpLw, s, t, im)
}

func (self *Code) lbu(t, s uint8, im int16) {
	self.iinst(inst.OpLbu, s, t, im)
}

func (self *Code) sw(t, s uint8, im int16) {
	self.iinst(inst.OpSw, s, t, im)
}

func (self *Code) lwStack(reg uint8, o StackObj) {
	self.lw(reg, regSP, o.Offset)
}

func (self *Code) swStack(reg uint8, o StackObj) {
	self.sw(reg, regSP, o.Offset)
}

func (self *Code) lbuStack(reg uint8, o StackObj) {
	self.lbu(reg, regSP, o.Offset)
}
