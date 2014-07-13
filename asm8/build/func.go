package build

import (
	"e8vm.net/e8/inst"
	"e8vm.net/leaf/tools/tok"
)

// Func is a continuous memory segment that saves executable instructions.
type Func struct {
	pos    *tok.Token
	name   string
	lines  []*Line
	labels map[string]int
}

func (f *Func) addLine(line *Line) {
	f.lines = append(f.lines, line)
}

func (f *Func) r3(fn, d, s, t uint8) {
	i := inst.Rinst(s, t, d, fn)
	f.addLine(newLine(i))
}

func (f *Func) r3r(fn, d, t, s uint8) {
	f.r3(fn, d, s, t)
}

func (f *Func) r3s(fn, d, t, sh uint8) {
	i := inst.RinstShamt(0, t, d, sh, fn)
	f.addLine(newLine(i))
}

func (f *Func) r3sSym(fn, d, t uint8, sh string) {
	i := inst.RinstShamt(0, t, d, 0, fn)
	f.addLine(newLineSym(i, sh))
}

func (f *Func) jump(op uint8, sym string) {
	i := inst.Jinst(op, 0)
	f.addLine(newLineSym(i, sym))
}
