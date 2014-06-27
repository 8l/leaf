package types

import (
	"bytes"
)

type Func struct {
	Ret  Type
	Args []Type
}

func NewFunc(r Type, args ...Type) *Func {
	ret := new(Func)
	ret.Ret = r
	ret.Args = make([]Type, len(args))
	copy(ret.Args, args)

	return ret
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
