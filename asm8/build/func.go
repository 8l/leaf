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

// op $d, $s, $t
func (f *Func) r3(fn, d, s, t uint8) {
	i := inst.Rinst(s, t, d, fn)
	f.addLine(newLine(i))
}

// op $d, $t, $s
func (f *Func) r3r(fn, d, t, s uint8) {
	f.r3(fn, d, s, t)
}

// op $d, $t, shamt
func (f *Func) r3s(fn, d, t, sh uint8) {
	i := inst.RinstShamt(0, t, d, sh, fn)
	f.addLine(newLine(i))
}

// op $d, $t, shamt
func (f *Func) r3sSym(fn, d, t uint8, sh string) {
	i := inst.RinstShamt(0, t, d, 0, fn)
	f.addLine(newLineSym(i, sh))
}

// op $s, $t, label
func (f *Func) i3sr(op, s, t uint8, lab string) {
	i := inst.Iinst(op, s, t, 0)
	f.addLine(newLineSym(i, lab))
}

// op $t, $s, im
func (f *Func) i3s(op, t, s uint8, im int16) {
	i := inst.Iinst(op, s, t, uint16(im))
	f.addLine(newLine(i))
}

// op $t, $s, im
func (f *Func) i3sSym(op, t, s uint8, im string) {
	i := inst.Iinst(op, s, t, 0)
	f.addLine(newLineSym(i, im))
}

// op Im
func (f *Func) jump(op uint8, sym string) {
	i := inst.Jinst(op, 0)
	f.addLine(newLineSym(i, sym))
}

// J appends a J that jumps to the symbol
func (f *Func) J(sym string) {
	f.jump(inst.OpJ, sym)
}

// Jal appends a Jal that jumps to the symbol
func (f *Func) Jal(sym string) {
	f.jump(inst.OpJal, sym)
}

// Beq appends a Beq that branches to the symbol
func (f *Func) Beq(s, t uint8, sym string) {
	f.i3sr(inst.OpBeq, s, t, sym)
}

// Bne appends a Bne that branches to the symbol
func (f *Func) Bne(s, t uint8, sym string) {
	f.i3sr(inst.OpBne, s, t, sym)
}

// Addi appends a Addi
func (f *Func) Addi(t, s uint8, im int16) {
	f.i3s(inst.OpAddi, t, s, im)
}

// AddiSym appends a Addi where the immediate is a symbol
func (f *Func) AddiSym(t, s uint8, im string) {
	f.i3sSym(inst.OpAddi, t, s, im)
}

// Slti appends a Slti
func (f *Func) Slti(t, s uint8, im int16) {
	f.i3s(inst.OpSlti, t, s, im)
}

// SltiSym appends a Slti where the immediate is a symbol
func (f *Func) SltiSym(t, s uint8, im string) {
	f.i3sSym(inst.OpSlti, t, s, im)
}
