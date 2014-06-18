package codegen

type Type interface {
	String() string
}

type BasicType struct {
	Name   string
	Actual int
}

func IsBasic(t Type) bool {
	_, ret := t.(*BasicType)
	return ret
}

func (t *BasicType) String() string {
	return t.Name
}

type PtrType struct {
	Type Type
}

func (t *PtrType) String() string {
	return "*" + t.Type.String()
}

type NamedType struct {
	Name string
	Type Type
}

func DeclareType(name string) *NamedType {
	ret := new(NamedType)
	ret.Name = name
	return ret
}

func (t *NamedType) String() string {
	return t.Name
}

const (
	Void = iota
	Int32
	Uint32
	Int8
	Uint8
)
