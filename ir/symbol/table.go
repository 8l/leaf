package symbol

type Table struct {
	stack []*Scope
}

func NewTable() *Table {
	ret := new(Table)
	return ret
}

func (self *Table) PushScope(s *Scope) {
	self.stack = append(self.stack, s)
}

func (self *Table) PopScope() {
	nstack := len(self.stack)
	if nstack == 0 {
		panic("bug")
	}

	self.stack = self.stack[:nstack-1]
}

func (self *Table) Get(n string) *Symbol {
	nstack := len(self.stack)
	for i := nstack - 1; i >= 0; i-- {
		s := self.stack[i]
		ret := s.Get(n)
		if ret != nil {
			return ret
		}
	}
	return nil
}
