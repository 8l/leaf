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

func (p *pos) newLine() {
	p.lineNo++
	p.lineOffset = 0
}

func (p *pos) syncTo(to *pos) {
	p.lineNo = to.lineNo
	p.lineOffset = to.lineOffset
}
