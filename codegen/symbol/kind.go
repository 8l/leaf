package symbol

type Kind int

const (
	Import Kind = iota
	Const
	Var
	Func
	Type
)
