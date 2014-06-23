package ir

type Inst struct {
	Op Op
	A  *Var
	B  *Var
	I  int32
}

type Op int

const (
	Push  Op = iota
	SpAdd    // add stack pointer
	Call
)
