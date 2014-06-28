package codegen

import (
	"e8vm.net/leaf/ir"
	"e8vm.net/leaf/parser/ast"
)

type task struct {
	fn   *ir.Func
	node *ast.Func
}
