package ir

import (
	"e8vm.net/leaf/leaf/ir/types"
)

type Func struct {
	name string
	t    *types.Func
	file *File
	code *Code
}

func (self *Func) Define() *Code {
	assert(self.code == nil)

	c := new(Code)
	c.table = self.file.newSymTable()
	c.retAddr = c.pushReg(regRet) // the return address

	// TODO: check the object size
	if self.t.Ret != nil {
		c.ret = c.fetchArg(int16(self.t.Ret.Size()))
	}

	narg := len(self.t.Args)
	c.args = make([]StackObj, narg)
	for i := narg - 1; i >= 0; i-- {
		arg := self.t.Args[i]
		c.args[i] = c.fetchArg(int16(arg.Size()))
	}

	self.code = c
	return c
}
