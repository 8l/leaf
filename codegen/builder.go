package codegen

import (
	"e8vm.net/leaf/ir"
	"e8vm.net/leaf/ir/symbol"
	"e8vm.net/leaf/ir/types"
	"e8vm.net/leaf/lexer"
	"e8vm.net/leaf/lexer/token"
	"e8vm.net/leaf/parser/ast"
)

type Gener struct {
	build *ir.Build
	pack  *ir.Package
	decls map[string]*decl // top level declares

	fileTrees []*ast.Program
	files     []*ir.File
}

func NewGener(path string, build *ir.Build) *Gener {
	ret := new(Gener)
	ret.build = build
	ret.pack = build.NewPackage(path)
	ret.decls = make(map[string]*decl)

	return ret
}

func (self *Gener) AddFile(p *ast.Program) {
	self.fileTrees = append(self.fileTrees, p)
}

func (self *Gener) Gen() {
	for _, ftree := range self.fileTrees {
		f := self.pack.NewFile(ftree.Filename)
		// TODO: register imports here
		self.files = append(self.files, f)
	}
	assert(len(self.files) == len(self.fileTrees))

	// first round, we register all the symbols
	for i, ftree := range self.fileTrees {
		self.symCheck(self.files[i], ftree)
	}

	// TODO: then we should resolve all the types and constants
	// in some proper order

	// third round, can now declare all the functions
	// at this point, all named types should be resolvable
	for i, ftree := range self.fileTrees {
		self.genCode(self.files[i], ftree)
	}
}

func (self *Gener) tryAddDecl(f *ir.File, newDecl *decl) *decl {
	name := newDecl.name
	old := self.decls[name]
	if old != nil {
		return old
	}

	self.decls[name] = newDecl
	return nil
}

func (self *Gener) errorf(t *lexer.Token, f string, args ...interface{}) {
	panic("todo")
}

func (self *Gener) symCheck(f *ir.File, prog *ast.Program) {
	for _, d := range prog.Decls {
		switch d := d.(type) {
		case *ast.Func:
			newDecl := &decl{
				class: symbol.Func,
				name:  d.Name,
				pos:   d.NameToken,
			}
			old := self.tryAddDecl(f, newDecl)
			if old != nil {
				self.errorf(newDecl.pos, "%s redeclared")
				self.errorf(old.pos, "   previously declared here")
			}
		default:
			panic("bug or todo")
		}
	}
}

func (self *Gener) genCode(file *ir.File, prog *ast.Program) {
	type funcGen struct {
		fn   *ir.Func
		node *ast.Func
	}

	// declare all the functions
	var funcs []*funcGen

	for _, d := range prog.Decls {
		f, isFunc := d.(*ast.Func)
		if !isFunc {
			continue
		}

		// build the func type
		assert(len(f.Args) == 0) // TODO
		assert(f.Ret == nil)
		var retType types.Type
		ft := types.NewFunc(retType)

		fn, _ := file.DeclNewFunc(f.Name, ft)
		funcs = append(funcs, &funcGen{fn, f})
	}

	// TODO: now declare all the variables
	// and also add anonymous init functions

	// BUG: need to decl are funcs first, and then generate the code

	// now generate all the func generate jobs
	for _, job := range funcs {
		self.genFunc(job.fn, job.node)
	}
}

func (self *Gener) genFunc(f *ir.Func, node *ast.Func) {
	code := f.Define() // build up the header

	code.EnterScope()
	// TODO: register the named args here

	self.genBlock(code, node.Block)
	code.ExitScope()

	code.Return() // always append a return at the end, just for safety
}

func (self *Gener) genBlock(code *ir.Code, b *ast.Block) {
	code.EnterScope()

	for _, stmt := range b.Stmts {
		self.genStmt(code, stmt)
	}

	code.ExitScope()
}

func (self *Gener) genStmt(code *ir.Code, s ast.Node) {
	switch s := s.(type) {
	default:
		panic("bug or todo")
	case *ast.EmptyStmt:
		return
	case *ast.ExprStmt:
		self.genExpr(code, s.Expr)
	}
}

func (self *Gener) genExpr(code *ir.Code, expr ast.Node) *obj {
	switch expr := expr.(type) {
	default:
		panic("bug or todo")
	case *ast.CallExpr:
		var args []*obj
		for _, arg := range expr.Args {
			o := self.genExpr(code, arg)
			args = append(args, o)
		}
		// TODO: push the args if nothing is error

		return nil // TODO:
	case *ast.Operand:
		return self.genOperand(code, expr)
	}
}

func (self *Gener) genOperand(code *ir.Code, op *ast.Operand) *obj {
	tok := op.Token

	switch tok.Token {
	default:
		panic("bug or todo")
	case token.Int:
		// parse the int and return an immediate obj
		panic("todo")

	case token.Char:
		// parse the char and return an immediate obj
		panic("todo")

	case token.Ident:
		return self.genIdent(code, tok.Lit)
	}
}

func (self *Gener) genIdent(code *ir.Code, ident string) *obj {
	panic("todo")
}
