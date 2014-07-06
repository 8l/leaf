package parser

import (
	"io"

	"e8vm.net/leaf/lexer"
	"e8vm.net/util/prt"
)

type trackNode interface{}

type tracker struct {
	root  *level
	stack []*level
}

func (t *tracker) add(n trackNode) {
	_, isLevel := n.(*level)
	_, isToken := n.(*lexer.Token)
	assert(isLevel || isToken)

	nlevel := len(t.stack)
	assert(nlevel > 0)
	top := t.stack[nlevel-1]
	top.add(n)
}

func (t *tracker) push(s string) {
	level := new(level)
	level.name = s

	if len(t.stack) == 0 {
		assert(t.root == nil)
		t.root = level
		t.stack = append(t.stack, level)
	} else {
		t.add(level)
		t.stack = append(t.stack, level)
	}
}

func (t *tracker) extend(s string) {
	assert(t.root != nil) // we cannot reduce when there is nothing
	assert(len(t.stack) != 0)

	level := new(level)
	level.name = s

	nlevel := len(t.stack)
	top := t.stack[nlevel-1]
	last := top.swapLast(level)

	level.add(last)
	t.stack = append(t.stack, level)
}

func (t *tracker) pop() trackNode {
	nlevel := len(t.stack)
	assert(nlevel > 0)
	top := t.stack[nlevel-1]
	t.stack = t.stack[:nlevel-1]
	return top
}

func (t *tracker) PrintTrack(out io.Writer) {
	if t.root == nil {
		return
	}

	p := prt.New(out)
	p.Indent = "    "

	t.root.PrintTo(p)
}
