package ast

import (
	"bytes"
	"fmt"
	"reflect"

	"e8vm.net/leaf/tools/prt"
)

func Print(p prt.Iface, n interface{}) {
	if n == nil {
		p.Print("! nil")
		return
	}

	switch n := n.(type) {
	case *Program:
		p.Printf("+ prog: %s", n.Filename)
		p.ShiftIn()
		for _, d := range n.Decls {
			Print(p, d)
		}
		p.ShiftOut()

	case *Func:
		p.Printf("+ func: %s", n.Name)
		Print(p, n.Block)

	case *Block:
		p.ShiftIn()
		for _, line := range n.Lines {
			Print(p, line)
		}
		p.ShiftOut()

	case *Line:
		if n.Label != nil {
			p.Printf("%s:", n.Label.Name)
		}

		if n.Inst != nil {
			p.Print("   ", n.Inst.String())
		}

	default:
		p.Printf("? %s: %s", reflect.TypeOf(n), n)
	}
}

func (i Inst) String() string {
	buf := new(bytes.Buffer)
	buf.WriteString(i.Op)

	for i, arg := range i.Args {
		if i > 0 {
			buf.WriteString(", ")
		} else {
			buf.WriteString(" ")
		}

		if arg.Reg != nil {
			fmt.Fprintf(buf, "$%s", arg.Reg.Lit)
		} else {
			if arg.Im != nil {
				fmt.Fprintf(buf, "%s", arg.Im.Lit)
			}
			if arg.Sym != nil {
				fmt.Fprintf(buf, "%s", arg.Sym.Lit)
			}

			if arg.AddrReg != nil {
				fmt.Fprintf(buf, "($%s)", arg.AddrReg.Lit)
			}
		}
	}

	return buf.String()
}
