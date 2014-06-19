package header

type Type interface{}

type Basic int

const (
	Void Basic = iota
	Int
	Uint
	Int32
	Uint32
	Int8
	Uint8
	Char
	Byte
)

type Pointer struct {
	Type Type
}

type Slice struct {
	Type Type
}

type Array struct {
	Type Type
	Size int
}

type Ref struct {
	Pacakge int // package index
	Symbol  int // symbol index
}

// struct is another thing
// it is basically description of a memory box
// and a set of named, typed offset
// sort of similar to function signature
