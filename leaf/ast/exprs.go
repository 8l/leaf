package ast

import (
	"e8vm.net/leaf/tools/tok"
)

type CallExpr struct {
	Func Node
	Args []Node

	Token *tok.Token
}

type Operand struct {
	Token *tok.Token
}
