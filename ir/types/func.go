package types

import (
	"bytes"
)

type Func struct {
	Ret  Type
	Args []Type
}

func (self *Func) Size() uint32 {
	return addrSize
}

func (self *Func) AddArg(t Type) {
	self.Args = append(self.Args, t)
}

func (self *Func) String() string {
	buf := new(bytes.Buffer)
	buf.WriteString("func (")
	for i, arg := range self.Args {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(arg.String())
	}
	buf.WriteString(")")

	if self.Ret != nil {
		buf.WriteString(" ")
		buf.WriteString(self.Ret.String())
	}

	return buf.String()
}
