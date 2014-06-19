package codegen

type namedType struct {
	name string
	typ  typ
}

var _ typ = new(namedType)
var _ symbol = new(namedType)

func (t *namedType) Size() uint32           { return t.typ.Size() }
func (t *namedType) String() string         { return t.name }
func (t *namedType) Name() string           { return t.name }
func defType(name string, t typ) *namedType { return &namedType{name, t} }
func declType(name string) *namedType       { return &namedType{name, nil} }
