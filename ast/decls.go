package ast

type Func struct {
	Name  string
	Args  []*FuncArg
	Ret   *FuncArg
	Block *Block
}

type FuncArg struct {
	Name string
	Type Node
}
