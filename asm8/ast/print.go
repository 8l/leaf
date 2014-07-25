package ast

import (
	"bytes"
	"fmt"
	"reflect"

	"e8vm.net/leaf/tools/prt"
)

// Print prints the AST with a printer.
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

	case *Var:
		p.Print(varString(n))
	default:
		p.Printf("? %s: %s", reflect.TypeOf(n), n)
	}
}

func varString(v *Var) string {
	ret := new(bytes.Buffer)
	fmt.Fprintf(ret, "+ var: %s", v.Name)
	if v.TypeToken != nil {
		if !v.IsArray {
			fmt.Fprintf(ret, " %s", v.Type)
		} else if v.SizeToken == nil {
			fmt.Fprintf(ret, " []%s", v.Type)
		} else {
			fmt.Fprintf(ret, " [%s]%s", v.SizeToken.Lit, v.Type)
		}
	}
	if v.InitValue != nil {
		fmt.Fprintf(ret, " = %s", v.InitValue.Lit)
	} else if v.InitValues != nil {
		fmt.Fprintf(ret, " = {...}")
	}

	return ret.String()
}

// String prints an instruction line using the printer.
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
