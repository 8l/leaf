package codegen

import (
	"e8vm.net/leaf/asm8/ast"
	"e8vm.net/leaf/asm8/build"
)

type funcTask struct {
	build *build.Func
	ast   *ast.Func

	lines []*lineTask
	start uint32
}

func newFuncTask(b *build.Func, a *ast.Func) *funcTask {
	ret := new(funcTask)
	ret.build = b
	ret.ast = a
	return ret
}
