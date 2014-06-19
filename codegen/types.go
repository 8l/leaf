package codegen

type Type interface {
	String() string
}

type basic int

type bacicType struct {
	Name string
	Type basic
}

func isBasic(t Type) bool {
	_, ret := t.(*bacicType)
	return ret
}

func (t *bacicType) String() string {
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

// basic
const (
	Void basic = iota
	Int32
	Uint32
	Int8
	Uint8
)
