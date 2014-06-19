package codegen

type ptrType struct {
	Type Type
}

func (t *ptrType) Size() uint32   { return 4 }
func (t *ptrType) String() string { return "*" + t.Type.String() }

var _ Type = new(ptrType)
