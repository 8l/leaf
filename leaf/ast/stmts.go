package ast

type Block struct {
	Stmts []Node
}

type EmptyStmt struct{}

type ExprStmt struct {
	Expr Node
}
