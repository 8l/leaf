package parser

import (
	"e8vm.net/leaf/lexer"
	"e8vm.net/util/prt"
)

type level struct {
	name string
	subs []trackNode
}

func (self *level) add(n trackNode) {
	self.subs = append(self.subs, n)
}

func (self *level) swapLast(n trackNode) trackNode {
	nsub := len(self.subs)

	assert(nsub > 0)

	ret := self.subs[nsub-1]
	self.subs[nsub-1] = n

	return ret
}

func (self *level) PrintTo(p prt.Iface) {
	p.Printf("+ %s:", self.name)
	p.ShiftIn()
	for _, sub := range self.subs {
		level, isLevel := sub.(*level)
		if isLevel {
			level.PrintTo(p)
		} else {
			tok := sub.(*lexer.Token)
			p.Print("- ", tok)
		}
	}
	p.ShiftOut()
}
