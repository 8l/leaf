package ir

import (
	"bytes"
)

/*
- our ir is type less
- ir are now a set of possiblly named, indexed boxes
*/

type Obj struct {
	Name string
	Addr uint32
	Buf *bytes.Buffer
}

func NewObj() *Obj {
	ret := new(Obj)
	ret.Buf = new(bytes.Buffer)
	return ret
}

func NewNamedObj(s string) *Obj {
	ret := NewObj()
	ret.Name = s
	return ret
}

func (self *Obj) Len() int {
	return self.Buf.Len()
}
