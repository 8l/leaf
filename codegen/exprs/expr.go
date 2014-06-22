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
		Base on which register
		- 0 for absolute
		- 30 for stack pointer
	*/
	Base uint8

	/*
		The address location that can be filled later.
	*/
	Fill *uint32

	/*
		The address location to use when Fill is nill.
	*/
	Offset uint32
}

type Type struct {
	Type types.Type
}

type err int

const Err err = 0
