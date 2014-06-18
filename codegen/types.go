package codegen

type Type interface{}

type BasicType struct {
	Name   string
	Actual int
}

type Pointer struct {
	Type Type
}

const (
	Void = iota
	Int32
	Uint32
	Int8
	Uint8
	Ptr
)
