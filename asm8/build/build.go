package build

import (
	"bytes"
	"io"

	"e8vm.net/e8/img"
	"e8vm.net/e8/mem"
	"e8vm.net/leaf/tools/tok"
)

// Build is an assembly build that consists of a set of functions, constants
// and variables.  For the first version, we do not plan to support linking.
type Build struct {
	funcs   []*Func
	funcMap map[string]*Func

	vars   []*Var
	varMap map[string]*Var
}

// NewBuild creates a new build.
func NewBuild() *Build {
	ret := new(Build)
	ret.funcMap = make(map[string]*Func)
	ret.varMap = make(map[string]*Var)
	return ret
}

// NewFunc creates a new function (code segment) in the build.
// If the function name is already defined, it will panic.
func (b *Build) NewFunc(name string, pos *tok.Token) *Func {
	if b.funcMap[name] != nil {
		panic("already defined")
	}

	ret := newFunc(name, pos)

	b.funcs = append(b.funcs, ret)
	b.funcMap[name] = ret

	return ret
}

// NewVar creates a new variable (data segment) in the build.
// If the variable name is already defined, it will panic.
func (b *Build) NewVar(name string, pos *tok.Token) *Var {
	if b.varMap[name] != nil {
		panic("already defined")
	}

	ret := newVar(name, pos)
	b.vars = append(b.vars, ret)
	b.varMap[name] = ret

	return ret
}

// Find returns the token for the declaration of the identifier.
// It can be used to check if the name is used to
// define a function, variable, or constant is
func (b *Build) Find(name string) (*tok.Token, SymType) {
	f := b.funcMap[name]
	if f != nil {
		return f.pos, SymFunc
	}

	v := b.varMap[name]
	if v != nil {
		return v.pos, SymVar
	}

	return nil, SymNone
}

// WriteImage writes out the program to an output stream
// in the form of an E8 image.
func (b *Build) WriteImage(out io.Writer) error {
	codeBuf := new(bytes.Buffer)
	for _, f := range b.funcs {
		f.emit(codeBuf)
	}

	dataBuf := new(bytes.Buffer)
	for _, v := range b.vars {
		v.emit(dataBuf)
	}

	w := img.NewWriter(out)
	e := w.Write(mem.SegCode, codeBuf.Bytes())
	if e != nil {
		return e
	}

	if dataBuf.Len() > 0 {
		e = w.Write(mem.SegHeap, dataBuf.Bytes())
		if e != nil {
			return e
		}
	}

	return nil
}
