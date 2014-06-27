package ir

type Func struct {
	name string
	t    TypeRef
	file *File
	code *Code
}
