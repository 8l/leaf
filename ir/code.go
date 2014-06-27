package ir

type Code struct {
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
