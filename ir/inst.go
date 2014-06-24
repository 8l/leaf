package ir

type Inst struct {
	Op Op
	A  *Var
	B  *Var
	I  int32
}

type Op int

const (
	Push Op = iota
	Link    // Load a block start address
	Call    // Call A
)
