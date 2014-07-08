package codegen

import (
	"e8vm.net/leaf/leaf/ast"
	"e8vm.net/leaf/leaf/ir"
)

type task struct {
	fn   *ir.Func
	node *ast.Func
}
