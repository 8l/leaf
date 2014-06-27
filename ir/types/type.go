package types

type Type interface {
	Size() uint32
	String() string // a single line string description
}

func TypeStr(t Type) string {
	if t == nil {
		return "<void>"
	}

	return t.String()
}
