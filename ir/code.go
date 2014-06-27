package ir

import (
	sym "e8vm.net/leaf/ir/symbol"
)

type Code struct {
	table *sym.Table
}

func (self *Code) Query(name string) *Ref {
	panic("todo")
}

func (self *Code) Push(v *Ref) {
}

func (self *Code) Call(f *Ref) *Ref {
	panic("todo")
}

func (self *Code) Return(f *Ref) {
}

func (self *Code) Number(i int64) *Ref {
	panic("todo")
}
