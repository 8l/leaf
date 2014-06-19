package codegen

type Type interface {
	Size() uint32
	String() string
}
