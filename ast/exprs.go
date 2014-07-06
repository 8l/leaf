package ast

import (
	"e8vm.net/leaf/lexer"
)

type CallExpr struct {
	Func Node
	Args []Node

	Token *lexer.Token
}

type Operand struct {
	Token *lexer.Token
}
