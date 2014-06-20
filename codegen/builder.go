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

// Returns IR code with symbol table
func (self *Builder) Build() (*Archive, []error) {
	self.archive = new(Archive)

	ret := func() (*Archive, []error) {
		return self.archive, self.errors
	}
	hasErrors := func() bool {
		return len(self.errors) > 0
	}

	for _, f := range self.files {
		self.defineTopDecls(f)
	}
	// might have duplicated declares
	if hasErrors() {
		return ret()
	}

	self.sortTopDecls() // sort decl building order
	// might have circular delcare dependencies
	if hasErrors() {
		return ret()
	}

	self.buildImports() // build the import symbols for each file context
	// might have missing imports
	if hasErrors() {
		return ret()
	}

	self.buildConstsAndTypes() // build the declarations (consts, types)
	if hasErrors() {
		return ret()
	}

	self.buildFuncSigs()
	if hasErrors() {
		return ret()
	}

	self.buildFuncsAndVars() // build the function bodies
	if hasErrors() {
		return ret()
	}

	self.buildInit() // determine init sequence

	return ret()
}

func (self *Builder) sortTopDecls()        {}
func (self *Builder) buildImports()        {}
func (self *Builder) buildConstsAndTypes() {}
func (self *Builder) buildFuncSigs()       {}
func (self *Builder) buildFuncsAndVars()   {}
func (self *Builder) buildInit()           {}

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

func (self *Builder) buildFile(arch *Archive, src *ast.Program) {

}

/*
	STEP 1: scan and load the imports
	we will skip this step for 0.1, because 0.1 will only target for 1
	single paaaackage however, we do need to import the builtin
	package, the package 0, and add the symbol table of the package
	into the base of our namespace
*/

/*
	STEP 2: collect symbols
	for the symbols, we have a tricky case
	for example, the program might say:

		func int8() { }

	which redefines int8 into a function rather than a type this means
	we cannot resolve named types in function signatures yet in this
	step at this step, we can only say that int8 is a function in this
	package. that's it.

	STEP 3: resolve the symbols
	we don't have custom named types in 0.1, so we don't need to
	resolve named types here in fact, we only need to resolve the
	function signatures there will only be types in function
	signatures, so all symbols required should be already resolved at
	this point we can hence build the function calling stack structure,
	because we know the size of each type of the calling args we can
	now also output the binary code here in future versions, we need to
	first resolve the types and consts

	STEP 4: we can generate the symbol table (of func symbols)
	for feeding other packages.

	Symbols:
	- a type symbol points to a type
	  and a (resolved) type always knows its size
	- a func symbol points to a func signature
	  and a (resolved) func signature always knows its calling stack
	  it also points to the starting address of the function in the lib
	- a var symbol points to a (resolved) type
	- a const symbol points to a constant

	The lib will contain two parts:
	- the lib signature:
		- symbols
		- exported function signature
		- exported structs

	If we want to support dynamic libraries, then a library/package
	must have an interface, which says these are the symbols we
	are going to use.

	We can define that a library interface can only be exported functions.
	and when using the library, you can only call the functions defined
	in the exported signature. The compiler will do the checking.

	Or simpler, we can define special dummy interafce packages, with all
	the neccessary data structures and consts, but minimal exported functions
	and vars.

*/
