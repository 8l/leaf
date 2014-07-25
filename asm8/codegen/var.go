package codegen

import (
	"math"

	"e8vm.net/e8/mem"
	"e8vm.net/leaf/asm8/ast"
	"e8vm.net/leaf/asm8/build"
	"e8vm.net/leaf/asm8/lexer/tt"
)

func (g *Gen) varDecl(d *ast.Var) {
	if d.IsArray {
		g.arrayVarDecl(d)
	} else {
		g.simpleVarDecl(d)
	}
}

func (g *Gen) parseAutoVarType(d *ast.Var) string {
	if d.InitValue != nil {
		if d.InitValue.Is(tt.String) {
			return "str"
		} else if d.InitValue.Is(tt.Int) {
			return "i32"
		} else if d.InitValue.Is(tt.Float) {
			return "f64"
		} else if d.InitValue.Is(tt.Char) {
			return "i8"
		}

		g.errorf(d.InitValue, "failed on detecting var type")
		return "err"
	} else if d.InitValues != nil {
		if len(d.InitValues) == 0 {
			// empty slice
			return "i32"
		}
		first := d.InitValues[0]
		if first.Is(tt.Int) {
			return "i32"
		} else if first.Is(tt.Float) {
			return "f64"
		} else if first.Is(tt.Char) {
			return "i8"
		}

		g.errorf(first, "failed on detecting var type in array")
		return "err"
	}

	return "i32" // no value, use i32 for default
}

func (g *Gen) parseVarType(d *ast.Var) string {
	if d.TypeToken == nil {
		// type is missing; auto-detect the var type
		return g.parseAutoVarType(d)
	}

	switch d.Type {
	case "string":
		return "str"
	case "int", "int32", "i32":
		return "i32"
	case "uint", "uint32", "u32":
		return "u32"
	case "uint8", "byte", "u8":
		return "u8"
	case "int8", "char", "i8":
		return "i8"
	case "float", "float64", "f64":
		return "f64"
	default:
		g.errorf(d.TypeToken, "unknown type %q", d.Type)
		return "err"
	}
}

func (g *Gen) simpleVarDecl(v *ast.Var) {
	name := v.Name
	pos, typ := g.build.Find(name)
	if pos != nil {
		g.errorf(v.NameToken, "%q already defined as a %s", name, typ)
		g.errorf(pos, "   %q previously defined here", name)
		return
	}

	t := g.parseVarType(v)
	if t == "err" {
		// the proper error is already reported
		return
	}

	switch t {
	case "str":
		g.declStrVar(v)
	case "i32":
		g.declIntVar(v, 4, true)
	case "u32":
		g.declIntVar(v, 4, false)
	case "i8":
		g.declIntVar(v, 1, true)
	case "u8":
		g.declIntVar(v, 1, false)
	case "f64":
		g.declFloatVar(v)
	default:
		panic("bug")
	}
}

func (g *Gen) declStrVar(v *ast.Var) {
	if v.InitValues != nil {
		g.errorf(v.NameToken, "need to init with a string, got an array")
		return
	}

	if v.InitValue == nil {
		g.errorf(v.NameToken, "string requires an init value")
		return
	}

	nv := g.newVar(v)
	value, e := parseStr(v.InitValue.Lit)
	if e != nil {
		g.errorf(v.InitValue, "invalid string")
		return
	}

	nv.WriteString(value)
}

func (g *Gen) newVar(v *ast.Var) *build.Var {
	name := v.Name
	ret := g.build.NewVar(name, v.NameToken)

	task := newVarTask(ret, v)
	g.vars = append(g.vars, task)
	g.varMap[name] = task

	return ret
}

func (g *Gen) declIntVar(v *ast.Var, unit int, signed bool) {
	if v.IsArray {
		// var count = 1
		panic("todo")
		return
	}

	nv := g.newVar(v)
	value := int64(0)
	nv.Align(uint32(unit))

	if v.InitValue != nil {
		var e error
		value, e = parseInt(v.InitValue.Lit)
		if e != nil {
			g.errorf(v.InitValue, "invalid init integer value")
			return
		}
	}

	switch {
	case unit == 1 && signed:
		if value >= math.MinInt8 && value <= math.MaxInt8 {
			nv.WriteByte(uint8(value))
		} else {
			g.errorf(v.InitValue, "uint8 init value out of range")
		}
	case unit == 1 && !signed:
		if value >= 0 && value <= math.MaxUint8 {
			nv.WriteByte(uint8(int8(value)))
		} else {
			g.errorf(v.InitValue, "int8 init value out of range")
		}
	case unit == 4 && signed:
		if value >= math.MinInt32 && value <= math.MaxInt32 {
			nv.WriteByte(uint8(value))
			nv.WriteByte(uint8(value >> 8))
			nv.WriteByte(uint8(value >> 16))
			nv.WriteByte(uint8(value >> 32))
		} else {
			g.errorf(v.InitValue, "int32 init value out of range")
		}
	case unit == 4 && !signed:
		if value >= 0 && value <= math.MaxUint32 {
			v := uint32(int32(value))
			nv.WriteByte(uint8(v))
			nv.WriteByte(uint8(v >> 8))
			nv.WriteByte(uint8(v >> 16))
			nv.WriteByte(uint8(v >> 32))
		} else {
			g.errorf(v.InitValue, "uint32 init value out of range")
		}
	default:
		panic("bug")
	}
}

func (g *Gen) declFloatVar(v *ast.Var) {
	panic("todo")
}

func (g *Gen) arrayVarDecl(v *ast.Var) {
	panic("todo")
}

func (g *Gen) layoutVars() {
	var offset uint32
	for _, v := range g.vars {
		offset = v.build.AlignStart(offset)
		v.start = mem.SegHeap + offset
		n := v.build.Len()
		if n > int(mem.SegSize) {
			g.errorf(v.ast.NameToken, "variable too large")
			return
		}

		offset += uint32(n)
		if offset > mem.SegSize {
			g.errorf(g.prog.EOFToken, "out of heap space")
			return
		}
	}
}

func (g *Gen) genVar(v *varTask) {

}
