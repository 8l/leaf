package codegen

import (
	"e8vm.net/leaf/asm8/ast"
	"e8vm.net/leaf/asm8/build"
)

type varTask struct {
	build *build.Var
	ast   *ast.Var

	start uint32
}

func newVarTask(b *build.Var, a *ast.Var) *varTask {
	ret := new(varTask)
	ret.build = b
	ret.ast = a
	return ret
}
