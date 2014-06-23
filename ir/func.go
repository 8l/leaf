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

func (self *Func) Link(link interface{}) *Var {
	ret := self.newVar()
	ret.Link = link
	return ret
}

func (self *Func) Push(vars ...*Var) {

}

/*
	func main:
		// no arg
		// no ret
		<0> = 42
		<1> = link "<builtin>.printInt"
		push <0>
		call <1>

	and from a building perspective

	f := NewFunc()
	_0 := f.Const(42)
	_1 := f.Link(printInt)
	f.Push(_0)
	f.Call(_1)

*/
