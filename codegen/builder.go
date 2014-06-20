package codegen

import (
	"fmt"

	"e8vm.net/leaf/ast"
	"e8vm.net/leaf/lexer"
)

type Builder struct {
	packName string
	files    []*ast.Program
	table    *symTable
	errors   []error
	archive  *Archive
}

func NewBuilder(name string) *Builder {
	ret := new(Builder)
	ret.packName = name
	ret.table = newSymTable()

	return ret
}

func (self *Builder) AddSource(src *ast.Program) {
	self.files = append(self.files, src)
}

func (self *Builder) errorf(t *lexer.Token, f string, args ...interface{}) {
	e := lexer.MakeError(t, fmt.Errorf(f, args...))
	self.errors = append(self.errors, e)
}

func (self *Builder) hasError() bool {
	return len(self.errors) > 0
}

// Returns IR code with symbol table
func (self *Builder) Build() (*Archive, []error) {
	self.archive = new(Archive)

	self.declare()
	// self.hookImports()
	self.buildDependency()
	self.buildSymbols()
	self.buildFunctions()

	return self.archive, self.errors
}

func (self *Builder) buildDependency() {
	if self.hasError() {
		return
	}
}

func (self *Builder) buildSymbols() {
	if self.hasError() {
		return
	}
}

func (self *Builder) buildFunctions() {
	if self.hasError() {
		return
	}
}

func (self *Builder) _declare(decl ast.Node) {
	switch decl := decl.(type) {
	case *ast.Func:
		declared := self.table.DeclTop(decl.Name, symFunc)
		if declared != nil {
			self.errorf(decl.Pos,
				"%q already declared as a %s",
				decl.Name, declared.kind,
			)
		}
	default:
		panic("bug: unknown decl in ast")
	}
}

func (self *Builder) declare() {
	if self.hasError() {
		return
	}

	for _, f := range self.files {
		for _, decl := range f.Decls {
			self._declare(decl)
		}
	}
}
