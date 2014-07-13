package codegen

import (
	"e8vm.net/leaf/asm8/ast"
	"e8vm.net/leaf/asm8/build"
)

// Gen generates the code into a build.
type Gen struct {
	build *build.Build
	prog  *ast.Program

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

func (g *Gen) declare(d ast.Decl) {
	switch d.(type) {
	case *ast.Const:
		panic("todo")
	case *ast.Var:
		panic("todo")
	case *ast.Func:
		panic("todo")
	}
}
