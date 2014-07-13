package codegen

import (
	"e8vm.net/leaf/asm8/ast"
	"e8vm.net/leaf/asm8/build"
	"e8vm.net/leaf/tools/tok"
)

// Gen generates the code into a build.
type Gen struct {
	build *build.Build
	prog  *ast.Program
	funcs []*build.Func

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
	for _, d := range g.prog.Decls {
		g.declare(d)
	}

	if len(g.errors) > 0 {
		return g.errors
	}

	return g.errors
}

func (g *Gen) errorf(pos *tok.Token, f string, args ...interface{}) {
	panic("todo")
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
		g.funcs = append(g.funcs, f)
	case *ast.Const:
		panic("todo")
	case *ast.Var:
		panic("todo")
	}
}
