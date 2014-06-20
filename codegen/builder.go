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

func (self *Builder) hasErrors() bool {
	return len(self.errors) > 0
}

// Returns IR code with symbol table
func (self *Builder) Build() (*Archive, []error) {
	self.archive = new(Archive)

	for _, f := range self.files {
		self.defineTopDecls(f)
	}
	if self.hasErrors() {
		return self.archive, self.errors
	}

	return self.archive, self.errors
}

func (self *Builder) errorf(t *lexer.Token, f string, args ...interface{}) {
	e := lexer.MakeError(t, fmt.Errorf(f, args...))
	self.errors = append(self.errors, e)
}

func (self *Builder) defineTopDecls(src *ast.Program) {
	// get a spot
	for _, decl := range src.Decls {
		switch decl := decl.(type) {
		case *ast.Func:
			defined := self.table.DeclTop(decl.Name, symFunc)
			if defined != nil {
				self.errorf(decl.Pos,
					"%q already declared as a %s",
					decl.Name, defined.kind,
				)
			}
		default:
			panic("bug: unknown decl in ast")
		}
	}

	if len(self.errors) > 0 {
		return
	}

	// TODO: sort the decls here in resolving order

	// for v0.1, we only have functions, so we can resolve in
	// whatever order we want

	// first, we resovle the function signatures for the functions
	// TODO:

	// now we can generate the function body for each function
	// TODO:
}
