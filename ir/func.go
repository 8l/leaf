package ir

type Func struct {
	vars  []*Var
	insts []*Inst
}

func NewFunc() *Func {
	ret := new(Func)
	return ret
}

func (self *Func) newVar() *Var {
	ret := new(Var)
	ret.Index = len(self.vars)
	ret.Size = 4
	self.vars = append(self.vars, ret)

	return ret
}

func (self *Func) newInst() *Inst {
	ret := new(Inst)
	self.insts = append(self.insts, ret)
	return ret
}

func (self *Func) Const(i uint32) *Var {
	ret := self.newVar()
	ret.Size = 4
	ret.IsConst = true
	ret.Value = i
	return ret
}

func (self *Func) Push(vars ...*Var) {

}
