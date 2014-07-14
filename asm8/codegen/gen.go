package codegen

import (
	"fmt"

	"e8vm.net/leaf/asm8/ast"
	"e8vm.net/leaf/asm8/build"
	"e8vm.net/leaf/tools/comperr"
	"e8vm.net/leaf/tools/tok"
)

// Gen generates the code into a build.
type Gen struct {
	build *build.Build
	prog  *ast.Program
	funcs []*funcTask

	errors []error
}

// NewGen creates a new code generator that
// generates the code into a build.
func NewGen(b *build.Build, p *ast.Program) *Gen {
	ret := new(Gen)
	ret.build = b
	ret.prog = p

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
		g.funcGen(f)
	}
	if len(g.errors) > 0 {
		return g.errors
	}

	// TODO: layout and output

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
		task := &funcTask{f, d}
		g.funcs = append(g.funcs, task)

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

// returns "" for invalid label
func parseLabel(arg *ast.Arg) string {
	if arg.Im != nil {
		return ""
	}
	if arg.Reg != nil {
		return ""
	}
	if arg.AddrReg != nil {
		return ""
	}
	if arg.Sym == nil {
		return ""
	}

	return arg.Sym.Lit
}

func (g *Gen) genJump(fn *build.Func, op string, args []*ast.Arg) bool {
	if len(args) != 1 {
		return false
	}

	lab := parseLabel(args[0])
	if lab == "" {
		return false
	}

	switch op {
	case "j":
		fn.J(lab)
	case "jal":
		fn.Jal(lab)
	default:
		panic("bug")
	}

	return true
}

func (g *Gen) funcGen(f *funcTask) {
	lines := f.ast.Block.Lines // ast node
	fn := f.build

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

		i := line.Inst // the instruction
		t := i.OpToken
		op := i.Op
		switch op {
		case "j", "jal":
			if !g.genJump(fn, op, i.Args) {
				g.synError(t, op)
			}

		default:
			g.errorf(t, "invalid op %q", op)
		}
	}
}
