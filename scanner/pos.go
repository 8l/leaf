package scanner

type pos struct {
	lineNo     int
	lineOffset int
}

func newPos() *pos {
	ret := new(pos)
	ret.lineNo = 1
	return ret
}

func (self *pos) Pos() uint32 {
	return uint32(self.lineNo)<<8 + uint32(self.lineOffset)
}

func (self *pos) NewLine() {
	self.lineNo++
	self.lineOffset = 0
}

func (self *pos) SyncTo(p *pos) {
	self.lineNo = p.lineNo
	self.lineOffset = p.lineOffset
}
