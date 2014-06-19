package codegen

import (
	"e8vm.net/leaf/ast"
)

type Builder struct {
	packName string
	files    []*ast.Program
	table    *symTable
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
func (self *Builder) Build() (output interface{}) {
	
	return nil
}
