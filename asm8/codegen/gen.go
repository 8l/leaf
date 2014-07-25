package codegen

import (
	"fmt"

	"e8vm.net/leaf/asm8/ast"
	"e8vm.net/leaf/asm8/build"
	"e8vm.net/leaf/tools/comperr"
	"e8vm.net/leaf/tools/tok"
)

// Gen generates the code into a build.
// TODO: we have a map in build and also a map in gen
// if we do not have a clear interface between gen and build,
// we should merge the two package together.
type Gen struct {
	build   *build.Build
	prog    *ast.Program
	funcs   []*funcTask
	funcMap map[string]*funcTask

	vars   []*varTask
	varMap map[string]*varTask

	errors []error
}

// NewGen creates a new code generator that
// generates the code into a build.
func NewGen(b *build.Build, p *ast.Program) *Gen {
	ret := new(Gen)
	ret.build = b
	ret.prog = p

	ret.funcMap = make(map[string]*funcTask)
	ret.varMap = make(map[string]*varTask)

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
		g.prepareFunc(f)
	}
	if len(g.errors) > 0 {
		return g.errors
	}

	g.layoutFuncs()
	if len(g.errors) > 0 {
		return g.errors
	}
	g.layoutVars()
	if len(g.errors) > 0 {
		return g.errors
	}

	for _, v := range g.vars {
		g.genVar(v)
	}
	if len(g.errors) > 0 {
		return g.errors
	}

	for _, f := range g.funcs {
		g.genFunc(f)
	}
	if len(g.errors) > 0 {
		return g.errors
	}

	return g.errors
}

// errorf creates an error that pins at the token
func (g *Gen) errorf(pos *tok.Token, f string, args ...interface{}) {
	e := comperr.New(pos, fmt.Errorf(f, args...))
	g.errors = append(g.errors, e)
}

func (g *Gen) declare(d ast.Decl) {
	switch d := d.(type) {
	case *ast.Func:
		g.declFunc(d)
	case *ast.Const:
		panic("todo")
	case *ast.Var:
		g.varDecl(d)
	}
}
