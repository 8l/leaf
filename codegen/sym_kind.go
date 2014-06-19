package codegen

type symKind int

const (
	symConst symKind = iota
	symType
	symFunc
	symVar
	symPackage
)
