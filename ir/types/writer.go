package types

import (
	"bytes"
)

func Marshall(t Type) []byte {
	w := newWriter()
	w.write(t)
	return w.Bytes()
}

type writer struct {
	*bytes.Buffer
}

func newWriter() *writer {
	ret := new(writer)
	ret.Buffer = new(bytes.Buffer)
	return ret
}

func (self *writer) writeStr(s string) {
	self.WriteString(s)
	self.WriteByte(0)
}

func (self *writer) write(t Type) {
	switch t := t.(type) {
	case nil:
		self.WriteByte(_nil)
	case Basic:
		self.WriteByte(_basic)
		self.WriteByte(byte(t))
	case *Pointer:
		self.WriteByte(_ptr)
		self.write(t.Type)
	case *Func:
		self.WriteByte(_func)
		self.write(t.Ret)
		for _, a := range t.Args {
			self.write(a)
		}
		self.WriteByte(_eol)
	case *Named:
		self.WriteByte(_named)
		self.writeStr(t.Path)
		self.writeStr(t.Name)
	default:
		panic("bug")
	}
}
