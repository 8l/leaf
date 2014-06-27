package ir

// the best data structure for this list is
// probably a hash linked list or a balanced binary tree
// anyways, we are just going to implement this in
// a stupid way
// TODO: use smarter algorithms
type typeList struct {
	types  []Type
	sorted []int
}

func newTypeList() *typeList {
	ret := new(typeList)

	return ret
}

func (self *typeList) add(t Type) {

}

func (self *typeList) find(t Type) {

}
