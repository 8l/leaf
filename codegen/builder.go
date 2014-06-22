package codegen

import (
	"fmt"
	"strconv"

	"e8vm.net/leaf/codegen/exprs"
	"e8vm.net/leaf/codegen/symbol"
	"e8vm.net/leaf/codegen/types"
	"e8vm.net/leaf/lexer"
	"e8vm.net/leaf/lexer/token"
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
	// this will be just type building on sized arrays and structs
	// for now, we can just move on

	self.gen()

	return self.object, self.errors
}

func (self *Builder) gen() {
	if self.hasError() {
		return
	}

	for _, decl := range self.prog.Decls {
		switch decl := decl.(type) {
		case *ast.Func:
			self.genFunc(decl)
		}
	}
}

func (self *Builder) genFunc(f *ast.Func) {
	// TODO: we should probably have an ir here
	// but for now, we will just write the assembly out directly
	// calling convention:
	// $30 is return address
	// $29 is stack pointer

	scope := symbol.NewScope()
	// define the named args here
	self.table.PushScope(scope)
	self.genBlock(f.Block)
	self.table.PopScope()
}

// Generates a block with its own scope
func (self *Builder) genBlock(b *ast.Block) {
	scope := symbol.NewScope()
	self.table.PushScope(scope)

	for _, stmt := range b.Stmts {
		self.genStmt(stmt)
	}

	self.table.PopScope()
}

// Generate a statement
func (self *Builder) genStmt(node ast.Node) {
	switch stmt := node.(type) {
	case *ast.EmptyStmt:
		return
	case *ast.Block:
		self.genBlock(stmt)
	case *ast.ExprStmt:
		self.genExpr(stmt.Expr)
	default:
		panic("bug")
	}
}

func (self *Builder) genOperand(op *ast.Operand) exprs.Expr {
	t := op.Token
	switch t.Token {
	case token.Ident:
		return self.genIdent(t.Lit)

	case token.Int:
		lit := t.Lit
		i, e := strconv.ParseInt(lit, 0, 64)
		if e != nil {
			self.errorf(t, "illegal integer; %s", e)
			return exprs.Err
		}

		ret := new(exprs.Int)
		ret.Value = i
		return ret

	default:
		panic("bug (or todo)")
	}
}

func (self *Builder) genIdent(s string) exprs.Expr {
	sym := self.table.Get(s)

	switch sym := sym.(type) {
	case *types.Named:
		ret := new(exprs.Type)
		ret.Type = sym
		return ret
	case *function:
		panic("todo") // this is a named function
	default:
		panic("bug (or todo)")
	}
}

func (self *Builder) genExpr(node ast.Node) exprs.Expr {
	switch expr := node.(type) {
	case *ast.CallExpr:
		args := make([]exprs.Expr, len(expr.Args))
		for i, arg := range expr.Args {
			args[i] = self.genExpr(arg)
		}
		// f := self.genExpr(expr.Func)

		// TODO: push the args on stack
		// and call the function
		// pop the args and save the result on stack
		return nil
	case *ast.Operand:
		return self.genOperand(expr)
	default:
		panic("bug")
	}
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

		var pre symbol.Symbol
		if s.Name() == "_" {
			// an anonymous top symbol
			ret = append(ret, s)
		} else {
			pre = self.scope.Register(s)
			if pre != nil {
				name := s.Name()
				self.errorf(s.Token(), "%q already declared", name)
				self.errorf(pre.Token(), "   %q previously declared here", name)
			} else {
				ret = append(ret, s)
			}
		}
	}

	return ret
}
