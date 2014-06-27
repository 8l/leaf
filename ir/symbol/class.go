package symbol

type Class int

const (
	Import Class = iota
	Const
	Var
	Func
	Type
)
