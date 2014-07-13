package ast

import (
	"e8vm.net/leaf/tools/tok"
)

// Decl defines a general declaration.
type Decl interface{}

// Program defines a source file.
type Program struct {
	Filename string
	Decls    []Decl
}

// BasicType defines the basic data types.
type BasicType int

// Basic data types with signed/unsigned and of all sizes.
const (
	U8 BasicType = iota
	I8
	U16
	I16
	U32
	I32
	F64
)

// Const defines a constant number.
type Const struct {
	Name  string
	Type  BasicType
	Value int
}

// Var defines a memory segment.
type Var struct {
	Name     string
	Size     int // 0 for auto, 1 for single
	Type     BasicType
	IsString bool

	Data    []int
	DataStr string
}

// Func defines a code segment.
type Func struct {
	Name  string
	Block *Block

	NameToken *tok.Token
}

// Block contains a set of instruction lines.
type Block struct {
	Lines []*Line
}

// Line is an assembly instruction line.
// It contains a label and an instruction.
// Both the label and the instruction are optional.
type Line struct {
	Label *Label
	Inst  *Inst
}

// Label marks a position in a code segment.
type Label struct {
	Name      string
	NameToken *tok.Token
}

// Inst is an instruction that will be translated into an
// assembly instruction
type Inst struct {
	Op      string
	OpToken *tok.Token
	Args    []*Arg
}

// Arg is an argument field for an instruction.
type Arg struct {
	Im      *tok.Token
	Reg     *tok.Token
	AddrReg *tok.Token
	Sym     *tok.Token
}