package codegen

import (
	"fmt"
	"math"

	"e8vm.net/e8/inst"
	"e8vm.net/e8/mem"
	"e8vm.net/leaf/asm8/ast"
	"e8vm.net/leaf/asm8/build"
	"e8vm.net/leaf/tools/comperr"
	"e8vm.net/leaf/tools/tok"
)

// Gen generates the code into a build.
type Gen struct {
	build   *build.Build
	prog    *ast.Program
	funcs   []*funcTask
	funcMap map[string]*funcTask

	errors []error
}

// NewGen creates a new code generator that
// generates the code into a build.
func NewGen(b *build.Build, p *ast.Program) *Gen {
	ret := new(Gen)
	ret.build = b
	ret.prog = p

	ret.funcMap = make(map[string]*funcTask)

	return ret
}

// Gen performs the code generation.
func (g *Gen) Gen() []error {
	// first round, declare the stuff
	for _, d := range g.prog.Decls {
		g.declare(d)
	}
	if len(g.errors) > 0 {
		return g.errors
	}

	for _, f := range g.funcs {
		g.funcPrepare(f)
	}
	if len(g.errors) > 0 {
		return g.errors
	}

	// layout
	offset := uint32(0)
	for _, f := range g.funcs {
		f.start = mem.SegCode + offset
		nline := len(f.lines)
		if nline > int(mem.SegSize)/4 {
			g.errorf(f.ast.NameToken,
				"too many instructions in this function")
			return g.errors
		}

		offset += 4 * uint32(nline)
		if offset > mem.SegSize {
			g.errorf(g.prog.EOFToken, "output code too large")
		}
	}

	for _, f := range g.funcs {
		g.funcGen(f)
	}
	if len(g.errors) > 0 {
		return g.errors
	}

	return g.errors
}

// creates an error that pins at the token
func (g *Gen) errorf(pos *tok.Token, f string, args ...interface{}) {
	e := comperr.New(pos, fmt.Errorf(f, args...))
	g.errors = append(g.errors, e)
}

func (g *Gen) declare(d ast.Decl) {
	switch d := d.(type) {
	case *ast.Func:
		name := d.Name
		pos, typ := g.build.Find(name)
		if pos != nil {
			g.errorf(d.NameToken, "%q already defined as a %s", name, typ)
			g.errorf(pos, "  %q previously defined here", name)
			return
		}

		f := g.build.NewFunc(name, d.NameToken)
		task := newFuncTask(f, d)
		g.funcs = append(g.funcs, task)
		g.funcMap[name] = task

	case *ast.Const:
		panic("todo")
	case *ast.Var:
		panic("todo")
	}
}

func (g *Gen) synError(pos *tok.Token, op string) {
	e := comperr.New(pos, fmt.Errorf("syntax error for %s", op))
	g.errors = append(g.errors, e)
}

func parseLabel(arg *ast.Arg) (lab string, valid bool) {
	if arg.Im != nil {
		return
	}
	if arg.Reg != nil {
		return
	}
	if arg.AddrReg != nil {
		return
	}
	if arg.Sym == nil {
		return
	}

	return arg.Sym.Lit, true
}

func parseReg(arg *ast.Arg) (reg uint8, valid bool) {
	if arg.Im != nil || arg.AddrReg != nil || arg.Sym != nil {
		return
	}
	if arg.Reg == nil {
		return
	}

	i, e := parseInt(arg.Reg.Lit)
	if e != nil {
		return
	}
	if i < 0 || i >= inst.Nreg {
		return
	}

	return uint8(i), true
}

func parseImm(arg *ast.Arg) (imm int64, valid bool) {
	if arg.AddrReg != nil || arg.Sym != nil || arg.Reg != nil {
		return
	}
	if arg.Im == nil {
		return
	}

	i, e := parseInt(arg.Im.Lit)
	if e != nil {
		return
	}
	return i, true
}

func parseAddr(arg *ast.Arg) (r uint8, imm int64, valid bool) {
	if arg.Reg != nil || arg.Sym != nil {
		return
	}

	if arg.AddrReg == nil {
		if arg.Im == nil {
			return
		}

		i, e := parseInt(arg.Im.Lit)
		if e != nil {
			return
		}
		return 0, i, true
	}

	reg, e := parseInt(arg.AddrReg.Lit)
	if e != nil {
		return
	}
	if reg < 0 || reg >= inst.Nreg {
		return
	}
	r = uint8(reg)

	if arg.Im != nil {
		i, e := parseInt(arg.Im.Lit)
		if e != nil {
			return
		}
		return r, i, true
	}

	return r, 0, true
}

func (g *Gen) funcPrepare(f *funcTask) {
	lines := f.ast.Block.Lines // ast node
	fn := f.build

	// scan labels and build line tasks
	for _, line := range lines {
		if line.Label != nil {
			label := line.Label
			t := label.NameToken
			got := fn.FindLabel(label.Name)
			if got != nil {
				g.errorf(t, "label %q already defined", t.Lit)
				g.errorf(got, "   previously defined here")
			} else {
				fn.MarkLabel(label.NameToken)
			}
		}

		if line.Inst == nil {
			continue
		}

		bline := fn.NewLine()
		f.lines = append(f.lines, &lineTask{bline, line})
	}
}

func (g *Gen) funcGen(f *funcTask) {
	for i, task := range f.lines {
		g.lineGen(f, i, task)
	}
}

func jumpInRange(off int) bool {
	return off >= (-1<<25) && off < (1<<25)
}

func shamtInRange(sh int64) bool {
	return sh >= 0 && sh < 32
}

func branchInRange(off int) bool {
	return off >= (-1<<15) && off < (1<<15)
}

func imuInRange(i int64) bool {
	return i >= 0 && i <= math.MaxUint16
}

func imsInRange(i int64) bool {
	return i >= math.MinInt16 && i <= math.MaxInt16
}

func indexStr(i int) string {
	switch i {
	case 1:
		return "first"
	case 2:
		return "second"
	case 3:
		return "third"
	case 4:
		return "forth"
	default:
		panic("bug")
	}
}

func (g *Gen) invalidArg(t *tok.Token, op string, i int, expect string) {
	istr := indexStr(i + 1)
	article := "a"
	if expect == "address" {
		article = "an"
	}
	g.errorf(t, "invalid %s arg for %s, expect %s %s",
		istr, op, article, expect,
	)
}

func (g *Gen) errorFmt(t *tok.Token, op string, argfmt string) {
	g.errorf(t, "error format;  expect: %s %s", op, argfmt)
}

func (g *Gen) lineGen(f *funcTask, index int, task *lineTask) {
	fn := f.build
	op := task.ast.Inst.Op
	t := task.ast.Inst.OpToken
	args := task.ast.Inst.Args
	line := task.build

	switch op {
	case "j", "jal":
		// j/jal <label>
		if len(args) != 1 {
			g.errorFmt(t, op, "<label>")
			return
		}
		lab, valid := parseLabel(args[0])
		if !valid {
			g.invalidArg(t, op, 0, "label")
			return
		}

		local := fn.LocateLabel(lab)
		if local >= 0 {
			delta := local - (index + 1)
			if !jumpInRange(delta) {
				g.errorf(t, "jump out of range")
				return
			}

			d := int32(delta)
			if op == "j" {
				line.Inst = inst.Jinst(inst.OpJ, d)
			} else {
				// "jal"
				line.Inst = inst.Jinst(inst.OpJal, d)
			}
			return
		}

		target := g.funcMap[lab]
		if target == nil {
			g.errorf(t, "jump target %q not found", lab)
			return
		}

		here := int(f.start/4) + index
		there := int(target.start / 4)
		delta := there - (here + 1)
		if !jumpInRange(delta) {
			g.errorf(t, "jump out of range")
			return
		}

		d := int32(delta)
		line.Inst = inst.Jinst(inst.OpCode(op), d)

	case "bne", "beq":
		// bne/beq $s, $t, <label>
		if len(args) != 3 {
			g.errorFmt(t, op, "$s, $t, <label>")
			return
		}

		rs, valid := parseReg(args[0])
		if !valid {
			g.invalidArg(t, op, 0, "reg")
			return
		}
		rt, valid := parseReg(args[1])
		if !valid {
			g.invalidArg(t, op, 1, "reg")
			return
		}
		lab, valid := parseLabel(args[2])
		if !valid {
			g.invalidArg(t, op, 2, "label")
			return
		}

		local := fn.LocateLabel(lab)
		if local < 0 {
			// label not found
			g.errorf(t, "label %q not found", lab)
			return
		}

		delta := local - (index + 1)
		if !branchInRange(delta) {
			g.errorf(t, "branch out of range, try use a jump")
			return
		}

		d := uint16(int16(delta))
		line.Inst = inst.Iinst(inst.OpCode(op), rs, rt, d)

	case "add", "sub", "and", "or", "xor", "nor", "slt",
		"mul", "mulu", "div", "divu", "mod", "modu":
		if len(args) != 3 {
			g.errorFmt(t, op, "$d, $s, $t")
			return
		}

		rd, valid := parseReg(args[0])
		if !valid {
			g.invalidArg(t, op, 0, "reg")
			return
		}
		rs, valid := parseReg(args[1])
		if !valid {
			g.invalidArg(t, op, 1, "reg")
			return
		}
		rt, valid := parseReg(args[2])
		if !valid {
			g.invalidArg(t, op, 2, "reg")
			return
		}

		line.Inst = inst.Rinst(rs, rt, rd, inst.FunctCode(op))

	case "sllv", "srlv", "srav":
		if len(args) != 3 {
			g.errorFmt(t, op, "$d, $t, $s")
			return
		}

		rd, valid := parseReg(args[0])
		if !valid {
			g.invalidArg(t, op, 0, "reg")
			return
		}
		rt, valid := parseReg(args[1])
		if !valid {
			g.invalidArg(t, op, 1, "reg")
			return
		}
		rs, valid := parseReg(args[2])
		if !valid {
			g.invalidArg(t, op, 2, "reg")
			return
		}

		line.Inst = inst.Rinst(rs, rt, rd, inst.FunctCode(op))

	case "sll", "srl", "sra":
		if len(args) != 3 {
			g.errorFmt(t, op, "$d, $t, shamt")
			return
		}

		rd, valid := parseReg(args[0])
		if !valid {
			g.invalidArg(t, op, 0, "reg")
			return
		}
		rt, valid := parseReg(args[1])
		if !valid {
			g.invalidArg(t, op, 1, "reg")
			return
		}
		shamt, valid := parseImm(args[2])
		if !valid {
			g.invalidArg(t, op, 2, "shamt")
			return
		}
		if !shamtInRange(shamt) {
			g.errorf(t, "shamt out of range", op)
			return
		}

		sh := uint8(shamt)
		line.Inst = inst.RinstShamt(0, rt, rd, sh, inst.FunctCode(op))

	case "andi", "ori":
		if len(args) != 3 {
			g.errorFmt(t, op, "$t, $s, imm")
			return
		}

		rt, valid := parseReg(args[0])
		if !valid {
			g.invalidArg(t, op, 0, "reg")
			return
		}
		rs, valid := parseReg(args[1])
		if !valid {
			g.invalidArg(t, op, 1, "reg")
			return
		}
		// TODO: allow using constant ident here
		im, valid := parseImm(args[2])
		if !valid {
			g.invalidArg(t, op, 2, "immediate")
			return
		}
		if !imuInRange(im) {
			g.errorf(t, "unsigned immediate out of range", op)
			return
		}

		imu := uint16(im)
		line.Inst = inst.Iinst(inst.OpCode(op), rs, rt, imu)

	case "addi", "slti":
		if len(args) != 3 {
			g.errorFmt(t, op, "$t, $s, imm")
			return
		}

		rt, valid := parseReg(args[0])
		if !valid {
			g.invalidArg(t, op, 0, "reg")
			return
		}
		rs, valid := parseReg(args[1])
		if !valid {
			g.invalidArg(t, op, 1, "reg")
			return
		}
		// TODO: allow using constant ident here
		im, valid := parseImm(args[2])
		if !valid {
			g.invalidArg(t, op, 2, "immediate")
			return
		}
		if !imsInRange(im) {
			g.errorf(t, "signed immediate out of range", op)
			return
		}

		ims := uint16(int16(im))
		line.Inst = inst.Iinst(inst.OpCode(op), rs, rt, ims)

	case "lui":
		if len(args) != 2 {
			g.errorFmt(t, op, "$t, imm")
			return
		}

		rt, valid := parseReg(args[0])
		if !valid {
			g.invalidArg(t, op, 0, "reg")
			return
		}

		// TODO: allow using constant ident here
		im, valid := parseImm(args[1])
		if !valid {
			g.invalidArg(t, op, 1, "immediate")
			return
		}
		if !imuInRange(im) {
			g.errorf(t, "unsigned immediate out of range", op)
			return
		}

		imu := uint16(im)
		line.Inst = inst.Iinst(inst.OpCode(op), 0, rt, imu)

	case "lw", "lu", "lhu", "lb", "lbu", "sw", "sh", "sb":
		if len(args) != 2 {
			g.errorFmt(t, op, "$t, imm($s)")
			return
		}
		rt, valid := parseReg(args[0])
		if !valid {
			g.invalidArg(t, op, 0, "reg")
			return
		}
		rs, im, valid := parseAddr(args[1])
		if !valid {
			g.invalidArg(t, op, 1, "address")
			return
		}
		if !imsInRange(im) {
			g.errorf(t, "signed immediate out of range", op)
			return
		}
		ims := uint16(int16(im))
		line.Inst = inst.Iinst(inst.OpCode(op), rs, rt, ims)

	default:
		g.errorf(t, "unknown instruction op name %q", op)
	}
}
