package ir

import (
	"bytes"
	"io"

	"e8vm.net/e8/img"
	"e8vm.net/e8/inst"
)

type linker struct {
	codes   []*Code
	codeMap map[Sym]*Code
	syms    map[Sym]uint32
}

func newLinker() *linker {
	ret := new(linker)
	ret.codeMap = make(map[Sym]*Code)
	ret.syms = make(map[Sym]uint32)
	return ret
}

func (self *linker) addCode(s Sym, c *Code) {
	assert(self.codeMap[s] == nil)
	self.codeMap[s] = c
}

func writeInst(out io.Writer, i inst.Inst) {
	var buf [4]byte
	buf[0] = byte(i)
	buf[1] = byte(i >> 8)
	buf[2] = byte(i >> 16)
	buf[3] = byte(i >> 24)
	_, e := out.Write(buf[:])
	assert(e == nil)
}

func setImm(i inst.Inst, im uint16) inst.Inst {
	return inst.Inst(uint32(i) | uint32(im))
}

func (self *linker) link(out io.Writer) []error {
	ptr := uint32(0)

	for _, c := range self.codes {
		c.start = ptr
		ptr += c.Size()
	}

	assert(ptr < segSize)

	for s, c := range self.codeMap {
		self.syms[s] = c.start
	}

	buf := new(bytes.Buffer)
	for _, c := range self.codes {
		assert(uint32(buf.Len()) == c.start)

		for _, i := range c.insts {
			if i.sym != nil {
				v, found := self.syms[*i.sym]
				assert(found)

				op := i.inst.Op()
				switch op {
				default:
					assert(uint16(i.inst) == 0)
					i.inst = setImm(i.inst, uint16(v))
				case inst.OpLui:
					assert(uint16(i.inst) == 0)
					i.inst = setImm(i.inst, uint16(v>>16))
				case inst.OpJ, inst.OpJal:
					// TODO: check jump range
					i.inst = inst.Jinst(op, int32(v-ptr+4))
				case inst.OpRinst, inst.OpBne, inst.OpBeq:
					panic("bug")
				}
			}

			writeInst(buf, i.inst)
		}
	}

	// TODO: avoid a double memory copy here
	e := img.Write(out, codeStart, buf.Bytes())
	if e != nil {
		return []error{e}
	}

	return nil
}
