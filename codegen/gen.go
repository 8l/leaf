package codegen

import (
	"strconv"

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
	funcs     []*funcGen
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
		self.symDecl(self.files[i], ftree)
	}

	// TODO: then we should resolve all the types and constants
	// in some proper order

	// third round, can now declare all the functions
	// at this point, all named types should be resolvable
	for i, ftree := range self.fileTrees {
		self.funcDecl(self.files[i], ftree)
	}

	// and finally, generate all the function bodies
	for i, ftree := range self.fileTrees {
		self.funcGen(self.files[i], ftree)
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

func (self *Gener) errore(t *lexer.Token, e error) {
	self.errorf(t, "%s", e.Error())
}

func (self *Gener) symDecl(f *ir.File, prog *ast.Program) {
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

func (self *Gener) funcDecl(file *ir.File, prog *ast.Program) {
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
		self.funcs = append(self.funcs, &funcGen{fn, f})
	}

	// TODO: now declare all the variables
	// and also add anonymous init functions
}

func (self *Gener) funcGen(file *ir.File, prog *ast.Program) {
	// now generate all the func generate jobs
	for _, job := range self.funcs {
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
		f := self.genExpr(code, expr.Func) // evaluate the function first
		if f == nil {
			return nil
		}

		ft, isFunc := f.t.(*types.Func)
		if !isFunc {
			self.errorf(expr.Token, "calling on a non-function")
			return nil
		}

		if len(expr.Args) != len(ft.Args) {
			self.errorf(expr.Token, "wrong number of arguments")
			return nil
		}

		var args []*obj
		for i, arg := range expr.Args {
			o := self.genExpr(code, arg)
			if o == nil {
				return nil
			}
			if !types.Equals(o.t, ft.Args[i]) {
				self.errorf(expr.Token, "wrong argument type")
				return nil
			}
			args = append(args, o)
		}

		// TODO: push the ret first
		ret := voidObj

		for _, o := range args {
			code.Push(o.o) // now we can push the stuff for call
		}

		code.Call(f)

		var pops []ir.Obj
		for _, o := range args {
			pops = append(pops, o.o)
		}
		code.Pop(pops...)

		return ret
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
		// TODO: use real large integers
		i, e := strconv.ParseInt(tok.Lit, 0, 64)
		if e != nil {
			self.errore(tok, e)
			return nil
		}
		return &obj{ir.ConstNum(i), types.ConstNum}
	case token.Char:
		c, e := unquoteChar(tok.Lit)
		if e != nil {
			self.errore(tok, e)
			return nil
		}
		return &obj{ir.ConstInt(int64(c), types.Int8), types.Int8}
	case token.Ident:
		return self.genIdent(code, tok)
	}
}

func (self *Gener) genIdent(code *ir.Code, tok *lexer.Token) *obj {
	assert(tok.Token == token.Ident)

	o, t := code.Query(tok.Lit)
	if o == nil {
		self.errorf(tok, "%q undefined", tok.Lit)
		return nil
	}

	return &obj{o, t}
}
