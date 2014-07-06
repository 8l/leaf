package ir

import (
	"e8vm.net/leaf/ir/types"
	"e8vm.net/util/prt"
	"reflect"
)

func (self *Build) Print() {
	for _, p := range self.packs {
		p.Print()
	}
}

func (self *Package) Print() {
	// just for debugging
	p := prt.Stdout()
	p.Indent = "    "

	p.Println("package", self.path)
	p.ShiftIn()

	for _, t := range self.types {
		name := t.Name()
		obj := t.Obj()
		switch t := obj.(type) {
		case *types.Named:
			p.Printf("type %s %s", name, t.Type)
		case types.Basic:
			p.Printf("type %s <%s>", name, t.String())
		case *types.Pointer:
			assert(t.Type == nil)
			p.Printf("type %s", name)
		default:
			panic(reflect.TypeOf(obj))
		}
	}

	for _, f := range self.funcs {
		name := f.Name()
		obj := f.Obj()
		switch f := obj.(type) {
		case *Func:
			p.Printf("func %s%s", name, f.t.SigString())
			if f.code != nil {
				p.ShiftIn()
				printInsts(p, f.code)
				p.ShiftOut()
			}
		default:
			panic("bug")
		}
	}

	if len(self.files) > 0 {
		p.Print("<files>")
		p.ShiftIn()
		for _, f := range self.files {
			p.Print(f.name)
		}
		p.ShiftOut()
	}

	p.ShiftOut()
	p.Println()
}

func printInsts(p *prt.Printer, c *Code) {
	for _, in := range c.insts {
		if in.sym == nil {
			p.Printf("%s", in.inst.String())
		} else {
			p.Printf("%s // %q.%s",
				in.inst.String(),
				in.sym.Pack,
				in.sym.Name,
			)
		}
	}
}
