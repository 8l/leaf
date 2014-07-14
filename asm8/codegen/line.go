package codegen

import (
	"e8vm.net/leaf/asm8/ast"
	"e8vm.net/leaf/asm8/build"
)

type lineTask struct {
	build *build.Line
	ast   *ast.Line
}
