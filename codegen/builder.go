package codegen

import (
	"e8vm.net/leaf/ast"
)

type Builder struct {
	packName string
	files    []*ast.Program
}

func NewBuilder(name string) *Builder {
	ret := new(Builder)
	ret.packName = name

	return ret
}

func (self *Builder) AddSource(src *ast.Program) {
	self.files = append(self.files, src)
}

func (self *Builder) Resolve() {

}
