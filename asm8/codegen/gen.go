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
