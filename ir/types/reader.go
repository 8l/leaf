package types

import (
	"bytes"
)

func Unmarshall(bs []byte) Type {
	r := newReader(bs)
	return r.read()
}

type reader struct {
	*bytes.Buffer
}

func newReader(bs []byte) *reader {
	ret := new(reader)
	ret.Buffer = bytes.NewBuffer(bs)
	return ret
}

func (self *reader) b() byte {
	b, e := self.ReadByte()
	if e != nil {
		panic(e)
	}
	return b
}

func (self *reader) readString() string {
	buf := new(bytes.Buffer)
	for b := self.b(); b != 0; b = self.b() {
		buf.WriteByte(b)
	}
	return buf.String()
}

func (self *reader) read() Type {
	b := self.b()

	switch b {
	case _nil:
		return nil
	case _basic:
		return Basic(self.b())
	case _ptr:
		t := self.read()
		return NewPointer(t)
	case _func:
		ret := new(Func)
		ret.Ret = self.read()
		for t := self.read(); !IsEOL(t); t = self.read() {
			ret.AddArg(t)
		}
		return ret
	case _named:
		ret := new(Named)
		ret.Path = self.readString()
		ret.Name = self.readString()
		return ret
	default:
		panic("bug")
	}
}
