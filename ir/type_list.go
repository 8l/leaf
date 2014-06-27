package ir

type typeList struct {
	types  []Type
	sorted []int
}

func newTypeList() *typeList {
	ret := new(typeList)

	return ret
}
