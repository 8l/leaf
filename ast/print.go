package ast

import (
	"reflect"

	"e8vm.net/leaf/prt"
)

func Print(p prt.Iface, n Node) {
	if n == nil {
		p.Printf("! nil")
		return
	}

	switch n := n.(type) {
	case *Program:
		p.Printf("+ src: %s", n.Filename)
		p.ShiftIn()
		for _, d := range n.Decls {
			Print(p, d)
		}
		p.ShiftOut()

	case *Func:
		p.Printf("+ func: %s", n.Name)
		if len(n.Args) > 0 {
			p.Printf("  args:")
			p.ShiftIn()
			for _, a := range n.Args {
				Print(p, a)
			}
			p.ShiftOut()
		}
		if n.Ret != nil {
			p.Printf("  ret: %s", n.Ret)
		}
		Print(p, n.Block)

	case *Block:
		p.Print("  {")
		p.ShiftIn()
		for _, s := range n.Stmts {
			Print(p, s)
		}
		p.ShiftOut("  }")

	case *EmptyStmt:
		p.Print("+ <empty-stmt>")

	case *ExprStmt:
		p.Print("+ <expr-stmt>:")
		p.ShiftIn()
		Print(p, n.Expr)
		p.ShiftOut()

	case *CallExpr:
		p.Print("+ <call-expr>")
		p.Print("  func:")
		p.ShiftIn()
		Print(p, n.Func)
		p.ShiftOut()

		if len(n.Args) > 0 {
			p.Print("  args:")
			p.ShiftIn()
			for _, a := range n.Args {
				Print(p, a)
			}
			p.ShiftOut()
		}

	case *Operand:
		p.Printf("+ <operand>: %s", n.Token)

	default:
		p.Printf("? %s: %s", reflect.TypeOf(n), n)
	}
}
