package tree

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	left   = 0
	center = 1
	right  = 2
)

// Tree is regular btree maker
type tree struct {
	Index int
}

// newTree returns new Tree
func newTree() *tree {
	return &tree{Index: 0}
}

// Node is Btree Leaf
type Node struct {
	tree   *tree
	Parent *Node `json:"-"`
	Left   *Node `json:"left"`
	Center *Node `json:"center"`
	Right  *Node `json:"right"`
	Index  int   `json:"index"`
}

// newNode returns new node referred by Tree.index
func (t *tree) newNode(p *Node) *Node {
	t.Index = t.Index + 1
	return &Node{
		tree:   t,
		Parent: p,
		Index:  t.Index,
	}
}

// NewNode returns new node referred by Tree.Index
func (t *Node) NewNode() *Node {
	return t.tree.newNode(t)
}

// Depth returns dist to root
func (t *Node) Depth() int {
	current, result := t, 0

	for {
		if current.Parent == nil {
			break
		}
		current = current.Parent
		result++
	}
	return result
}

// Children returns slice [ Left, Center, Right ]
func (t *Node) Children() []*Node {
	result := make([]*Node, 3)
	result[left] = t.Left
	result[center] = t.Center
	result[right] = t.Right
	return result
}

// Dig returns self digged 1 generations
func (t *Node) Dig() []*Node {
	t.Left = t.NewNode()
	t.Center = t.NewNode()
	t.Right = t.NewNode()
	return t.Children()
}

// Dig returns digged tree combination
func Dig(nodes ...*Node) []*Node {
	result := make([]*Node, len(nodes)*3)
	for i, node := range nodes {
		node.Dig()
		result[i*3+left] = node.Left
		result[i*3+center] = node.Center
		result[i*3+right] = node.Right
	}
	return result
}

// New returns root node structured regularlly
func New(depth int) *Node {
	tree := newTree()
	node := tree.newNode(nil)
	result := Dig(node)

	i := 1
	for {
		result = Dig(result...)
		i++

		if i >= depth {
			break
		}
	}

	return node
}

// ServeHTTP handles net/http interface.
func (t *Node) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(t)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	} else {
		w.Write(data)
	}
}

// Find returns the node having argument index
func (t *Node) Find(index int) *Node {
	if t.Index == index {
		return t
	}

	var result *Node
	for _, n := range t.Children() {
		if n != nil {
			result = n.Find(index)
		}
		if result != nil {
			break
		}
	}
	return result
}
