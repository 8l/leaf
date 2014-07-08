package codegen

import (
	"e8vm.net/leaf/leaf/ast"
	"e8vm.net/leaf/leaf/ir"
)

func (self *Gen) genBlock(code *ir.Code, b *ast.Block) {
	code.EnterScope()

	for _, stmt := range b.Stmts {
		self.genStmt(code, stmt)
	}

	code.ExitScope()
}

func (self *Gen) genStmt(code *ir.Code, s ast.Node) {
	switch s := s.(type) {
	default:
		panic("bug or todo")
	case *ast.EmptyStmt:
		return
	case *ast.ExprStmt:
		self.genExpr(code, s.Expr)
	}
}
