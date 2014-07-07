package codegen

import (
	"fmt"

	"e8vm.net/leaf/ast"
	"e8vm.net/leaf/ir"
	"e8vm.net/leaf/ir/symbol"
	"e8vm.net/leaf/ir/types"
	"e8vm.net/util/comperr"
	"e8vm.net/util/tok"
)

type Gen struct {
	build *ir.Build
	pack  *ir.Package
	decls map[string]*decl // top level declares

	fileTrees []*ast.Program
	files     []*ir.File
	tasks     []*task
	errors    []error
}

func NewGen(path string, build *ir.Build) *Gen {
	ret := new(Gen)
	ret.build = build
	ret.pack = build.NewPackage(path)
	ret.decls = make(map[string]*decl)

	return ret
}

func (self *Gen) AddFile(p *ast.Program) {
	self.fileTrees = append(self.fileTrees, p)
}

func (self *Gen) Gen() []error {
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
	if len(self.errors) > 0 {
		return self.errors
	}

	// TODO: second round, we resolve all the types and constants
	// in some proper order

	// third round, can now declare all the functions
	// at this point, all named types should be resolvable
	for i, ftree := range self.fileTrees {
		self.funcDecl(self.files[i], ftree)
	}
	if len(self.errors) > 0 {
		return self.errors
	}

	// and finally, generate all the function bodies
	for i, ftree := range self.fileTrees {
		self.funcGen(self.files[i], ftree)
	}
	if len(self.errors) > 0 {
		return self.errors
	}

	return self.errors
}

func (self *Gen) tryAddDecl(f *ir.File, newDecl *decl) *decl {
	name := newDecl.name
	old := self.decls[name]
	if old != nil {
		return old
	}

	self.decls[name] = newDecl
	return nil
}

func (self *Gen) errorf(t *tok.Token, f string, args ...interface{}) {
	e := comperr.New(t, fmt.Errorf(f, args...))
	self.errors = append(self.errors, e)
}

func (self *Gen) errore(t *tok.Token, e error) {
	self.errorf(t, "%s", e.Error())
}

func (self *Gen) symDecl(f *ir.File, prog *ast.Program) {
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
				self.errorf(newDecl.pos, "%q redeclared", newDecl.name)
				self.errorf(old.pos, "    previously declared here")
			}
		default:
			panic("bug or todo")
		}
	}
}

func (self *Gen) funcDecl(file *ir.File, prog *ast.Program) {
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
		self.tasks = append(self.tasks, &task{fn, f})
	}

	// TODO: now declare all the variables
	// and also add anonymous init functions
}

func (self *Gen) funcGen(file *ir.File, prog *ast.Program) {
	// now generate all the func generate jobs
	for _, job := range self.tasks {
		self.defineFunc(job.fn, job.node)
	}
}

func (self *Gen) defineFunc(f *ir.Func, node *ast.Func) {
	code := f.Define() // build up the header

	code.EnterScope()
	// TODO: register the named args here

	self.genBlock(code, node.Block)
	code.ExitScope()

	code.Return() // always append a return at the end, just for safety
}
