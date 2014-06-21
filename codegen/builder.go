package codegen

import (
	"fmt"

	"e8vm.net/leaf/codegen/symbol"
	"e8vm.net/leaf/lexer"
	"e8vm.net/leaf/parser/ast"
)

type Builder struct {
	prog   *ast.Program
	scope  *symbol.Scope // package level symbols
	table  *symbol.Table
	syms   []symbol.Symbol
	errors []error
	object *Object
}

func NewBuilder(p *ast.Program) *Builder {
	ret := new(Builder)
	ret.prog = p
	ret.scope = symbol.NewScope()
	ret.table = symbol.NewTable()
	// ret.table.Push(builtin)
	// ret.table.Push(ret.scope)

	return ret
}

func (self *Builder) errorf(t *lexer.Token, f string, args ...interface{}) {
	e := lexer.MakeError(t, fmt.Errorf(f, args...))
	self.errors = append(self.errors, e)
}

func (self *Builder) hasError() bool {
	return len(self.errors) > 0
}

// Returns IR code with symbol table
func (self *Builder) Build() (*Object, []error) {
	self.object = new(Object)

	self.register()

	self.syms = self.scope.List() // get a copy of all symbols

	return self.object, self.errors
}

func (self *Builder) register() {
	if self.hasError() {
		return
	}

	for _, decl := range self.prog.Decls {
		var s symbol.Symbol

		switch decl := decl.(type) {
		case *ast.Func:
			s = newFunc(decl.NameToken)
		default:
			panic("bug: unknown decl in ast")
		}

		pre := self.scope.Register(s)
		if pre != nil {
			name := s.Name()
			self.errorf(s.Token(), "%q already declared", name)
			self.errorf(pre.Token(), "   %q previously declared here", name)
		}
	}
}
