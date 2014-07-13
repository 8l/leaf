package codegen

import (
	"e8vm.net/leaf/asm8/ast"
	"e8vm.net/leaf/asm8/build"
)

type funcTask struct {
	build *build.Func
	ast   *ast.Func
}
