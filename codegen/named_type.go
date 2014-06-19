package codegen

type namedType struct {
	name string
	typ  typ
}

func (t *namedType) Size() uint32           { return t.typ.Size() }
func (t *namedType) String() string         { return t.name }
func defType(name string, t typ) *namedType { return &namedType{name, t} }
func declType(name string) *namedType       { return &namedType{name, nil} }
