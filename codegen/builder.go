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

	self.syms = self.register()

	// now that all the symbols are registered
	// this is only layout names
	// self.layout() // TODO: evaluate consts and types,
	// TODO: we also need to figure out the types of the vars
	// we could require that a variable must declare a type first here
	// rather than detecting the type based on the expression
	// and the function signatures

	self.implement() // function and init implementations

	return self.object, self.errors
}

func (self *Builder) implement() {

}

func (self *Builder) register() []symbol.Symbol {
	if self.hasError() {
		return nil
	}
	var ret []symbol.Symbol

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
		} else {
			ret = append(ret, s)
		}
	}

	return ret
}
