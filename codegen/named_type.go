package codegen

type namedType struct {
	Name string
	Type Type // when it is
}

func (t *namedType) Size() uint32   { return t.Type.Size() }
func (t *namedType) String() string { return t.Name }
