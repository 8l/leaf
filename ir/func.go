package ir

import (
	"e8vm.net/leaf/ir/types"
)

type Func struct {
	name string
	t    *types.Func
	file *File
	code *Code
}

func (self *Func) Define() *Code {
	assert(self.code == nil)

	ret := new(Code)
	ret.table = self.file.newSymTable()

	self.code = ret

	c := self.code
	c.ret = c.pushReg(regRet) // the return address

	return ret
}
