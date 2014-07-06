package codegen

import (
	"e8vm.net/leaf/ast"
	"e8vm.net/leaf/ir"
)

type task struct {
	fn   *ir.Func
	node *ast.Func
}
