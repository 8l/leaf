package codegen

import (
	"e8vm.net/e8/mem"
	"e8vm.net/leaf/asm8/ast"
)

func (g *Gen) layoutFuncs() {
	offset := uint32(0)
	for _, f := range g.funcs {
		f.start = mem.SegCode + offset
		nline := len(f.lines)
		if nline > int(mem.SegSize)/4 {
			g.errorf(f.ast.NameToken,
				"too many instructions in this function")
			return
		}

		offset += 4 * uint32(nline)
		if offset > mem.SegSize {
			g.errorf(g.prog.EOFToken, "output code too large")
			return
		}
	}
}

func (g *Gen) declFunc(d *ast.Func) {
	name := d.Name
	pos, typ := g.build.Find(name)
	if pos != nil {
		g.errorf(d.NameToken, "%q already defined as a %s", name, typ)
		g.errorf(pos, "  %q previously defined here", name)
		return
	}

	f := g.build.NewFunc(name, d.NameToken)
	task := newFuncTask(f, d)
	g.funcs = append(g.funcs, task)
	g.funcMap[name] = task
}

func (g *Gen) genFunc(f *funcTask) {
	for i, task := range f.lines {
		g.lineGen(f, i, task)
	}
}

// prepareFunc marks all the label positions and
// prepares all the line tasks for func gen,
func (g *Gen) prepareFunc(f *funcTask) {
	lines := f.ast.Block.Lines // ast node
	fn := f.build

	// scan labels and build line tasks
	for _, line := range lines {
		if line.Label != nil {
			label := line.Label
			t := label.NameToken
			got := fn.FindLabel(label.Name)
			if got != nil {
				g.errorf(t, "label %q already defined", t.Lit)
				g.errorf(got, "   previously defined here")
			} else {
				fn.MarkLabel(label.NameToken)
			}
		}

		if line.Inst == nil {
			continue
		}

		bline := fn.NewLine()
		f.lines = append(f.lines, &lineTask{bline, line})
	}
}
