package codegen

type typ interface {
	Size() uint32
	String() string
}
