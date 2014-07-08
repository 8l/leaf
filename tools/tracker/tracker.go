// Package tracker provides a tree builder.
package tracker

import (
	"io"

	"e8vm.net/leaf/tools/prt"
)

// Node is a general branch or leaf in the tracking tree.
type Node interface {
	String() string
}

// Tracker is a tree builder.
type Tracker struct {
	root  *Level
	stack []*Level
}

// New creates an empty tracker.
func New() *Tracker {
	return new(Tracker)
}

// Add appends a node to the top level on the stack
func (t *Tracker) Add(n Node) {
	nlevel := len(t.stack)
	if nlevel == 0 {
		panic("add to root")
	}

	t.stack[nlevel-1].add(n)
}

// Push pushes a new level onto the stack.
func (t *Tracker) Push(s string) {
	level := newLevel(s)

	if len(t.stack) == 0 {
		assert(t.root == nil)
		t.root = level
		t.stack = append(t.stack, level)
	} else {
		t.Add(level)
		t.stack = append(t.stack, level)
	}
}

// Extend extends the last node on the top level into a new
// level on the stack.
func (t *Tracker) Extend(s string) {
	if len(t.stack) == 0 {
		panic("no stuff to extend")
	}

	level := newLevel(s)

	nlevel := len(t.stack)
	last := t.stack[nlevel-1].swapLast(level)
	level.add(last)
	t.stack = append(t.stack, level)
}

// Pop pops the top level out of the stack.
func (t *Tracker) Pop() Node {
	nlevel := len(t.stack)
	if nlevel == 0 {
		panic("not stuff to pop")
	}

	top := t.stack[nlevel-1]
	t.stack = t.stack[:nlevel-1]
	return top
}

// Print prints the tree out in an indented format using
// the output stream.
func (t *Tracker) Print(out io.Writer) {
	if t.root == nil {
		return
	}

	p := prt.New(out)
	p.Indent = "    "
	t.root.PrintTo(p)
}
