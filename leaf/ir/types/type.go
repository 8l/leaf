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

func Equals(t1, t2 Type) bool {
	b1 := Marsh(t1)
	b2 := Marsh(t2)
	if len(b1) != len(b2) {
		return false
	}
	for i := range b1 {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
}
