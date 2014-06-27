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

func (self *Code) addi(t, s uint8, i uint16) {
	self.inst(inst.Iinst(inst.OpAddi, s, t, i))
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
	self.addi(reg, 0, uint16(i))
}

func (self *Code) loadiSym(reg uint8, sym *Sym) {
	self.luiSym(reg, sym)
	self.addiSym(reg, 0, sym)
}
