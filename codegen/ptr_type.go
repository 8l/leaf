package codegen

type ptrType struct {
	t typ
}

func (t *ptrType) Size() uint32   { return 4 }
func (t *ptrType) String() string { return "*" + t.t.String() }

var _ typ = new(ptrType)

func ptr(t typ) *ptrType { return &ptrType{t} }
