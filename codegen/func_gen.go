package codegen

import (
	"e8vm.net/leaf/ir"
	"e8vm.net/leaf/parser/ast"
)

type funcGen struct {
	fn   *ir.Func
	node *ast.Func
}
