package tracker

import (
	"e8vm.net/leaf/tools/prt"
)

// Level is a tracker level on the tracker stack.
type Level struct {
	Nodes []Node
	Name  string
}

func newLevel(name string) *Level {
	ret := new(Level)
	ret.Name = name
	return ret
}

func (lvl *Level) add(n Node) {
	lvl.Nodes = append(lvl.Nodes, n)
}

func (lvl *Level) swapLast(n Node) Node {
	nnode := len(lvl.Nodes)
	if nnode == 0 {
		panic("no node to swap")
	}

	ret := lvl.Nodes[nnode-1]
	lvl.Nodes[nnode-1] = n
	return ret
}

// PrintTo prints the level using the indent printer
func (lvl *Level) PrintTo(p prt.Iface) {
	p.Printf("+ %s:", lvl.Name)
	p.ShiftIn()

	for _, node := range lvl.Nodes {
		level, isLevel := node.(*Level)

		if isLevel {
			level.PrintTo(p)
		} else {
			p.Printf("- %s", node)
		}
	}
	p.ShiftOut()
}

// String returns the indented string presentation
// of this level's nodes and its sub-levels.
func (lvl *Level) String() string {
	return prt.String(lvl)
}
