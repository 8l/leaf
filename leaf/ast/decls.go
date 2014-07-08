package ast

import (
	"e8vm.net/leaf/tools/tok"
)

type Func struct {
	Name  string
	Args  []*FuncArg
	Ret   *FuncArg
	Block *Block

	NameToken *tok.Token
}

type FuncArg struct {
	Name string
	Type Node
}
