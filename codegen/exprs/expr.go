package exprs

import (
	"e8vm.net/leaf/codegen/types"
)

// expression result
// often a location in memory
// a stack location, or an absolute location
type Expr interface{}

type Int struct {
	Type  types.Type
	Value int64 // should be good for int values
}

type Addr struct {
	/*
		Type of the address. Must be one of those:
		- A pointer, where it saves the pointed address
		- A function, where it saves the start PC of the function
		- A structure, where it saves the start location of the struct
	*/
	Type types.Type

	/*
		Location base on which register
		- 0 for absolute
		- 30 for stack pointer
		- 31 for relative jumping (relative PC)
	*/
	LocBase uint8

	/*
		The address location to use when Fill is nill.
	*/
	Loc uint32

	ReadOnly bool
}

type Ref struct {
	Type types.Type
	ReadOnly bool
}

// can be used as a function for casting type into the target type
type TypeCast struct {
	Type types.Type
}

type err int

const Err err = 0
