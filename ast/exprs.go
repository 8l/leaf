package ast

import (
	"e8vm.net/util/tok"
)

type CallExpr struct {
	Func Node
	Args []Node

	Token *tok.Token
}

type Operand struct {
	Token *tok.Token
}
