package ast

import (
	"e8vm.net/leaf/lexer"
)

type Func struct {
	Name  string
	Args  []*FuncArg
	Ret   *FuncArg
	Block *Block

	Pos *lexer.Token
}

type FuncArg struct {
	Name string
	Type Node
}
