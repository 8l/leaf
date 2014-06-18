package tree

type Data interface {
	String() string
}

type Node struct {
	Data     Data
	Father   *Node
	Children []*Node
}

func NewNode(data Data) *Node {
	ret := new(Node)
	ret.Data = data
	return ret
}

func (self *Node) Add(n *Node) {
	self.Children = append(self.Children, n)
	n.Father = self
}

func (self *Node) AddTo(n *Node) {
	n.Add(self)
}
