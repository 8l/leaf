package types

import (
	"bytes"
	"io"
)

func Marsh(t Type) []byte {
	buf := new(bytes.Buffer)
	w := newWriter(buf)
	w.write(t)
	return buf.Bytes()
}

type writer struct {
	io.ByteWriter
}

func newWriter(w io.ByteWriter) *writer {
	ret := new(writer)
	ret.ByteWriter = w
	return ret
}

func (self *writer) writeStr(s string) {
	for _, b := range []byte(s) {
		self.WriteByte(b)
	}
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
