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

func (self *Code) instSym(i inst.Inst, s *Sym) {
	self.insts = append(self.insts, Inst{i, s})
}

func (self *Code) add(d, s, t uint8) {
	self.inst(inst.Rinst(s, t, d, inst.FnAdd))
}

func (self *Code) mov(d, s uint8) {
	self.add(d, s, 0)
}

func (self *Code) addi(t, s uint8, i int16) {
	self.inst(inst.Iinst(inst.OpAddi, s, t, uint16(i)))
}

func (self *Code) subi(t, s uint8, i int16) {
	self.addi(t, s, -i)
}

func (self *Code) addiSym(t, s uint8, sym *Sym) {
	self.instSym(inst.Iinst(inst.OpAddi, s, t, 0), sym)
}

func (self *Code) lui(t uint8, i uint16) {
	self.inst(inst.Iinst(inst.OpLui, 0, t, i))
}

func (self *Code) luiSym(t uint8, sym *Sym) {
	self.instSym(inst.Iinst(inst.OpLui, 0, t, 0), sym)
}

func (self *Code) loadi(reg uint8, i uint32) {
	self.lui(reg, uint16(i>>16))
	self.addi(reg, 0, int16(uint16(i)))
}

func (self *Code) loadiSym(reg uint8, sym *Sym) {
	self.luiSym(reg, sym)
	self.addiSym(reg, 0, sym)
}

func (self *Code) jSym(sym *Sym) {
	self.instSym(inst.Jinst(inst.OpJ, 0), sym)
}

func (self *Code) jalSym(sym *Sym) {
	self.instSym(inst.Jinst(inst.OpJal, 0), sym)
}

func (self *Code) jr(reg uint8) {
	self.mov(regPC, reg)
}

func (self *Code) lw(t, s uint8, im int16) {
	self.inst(inst.Iinst(inst.OpLw, s, t, uint16(im)))
}

func (self *Code) sw(t, s uint8, im int16) {
	self.inst(inst.Iinst(inst.OpSw, s, t, uint16(im)))
}

func (self *Code) lstack(reg uint8, o StackObj) {
	self.lw(reg, regSP, o.Offset)
}

func (self *Code) sstack(reg uint8, o StackObj) {
	self.sw(reg, regSP, o.Offset)
}
