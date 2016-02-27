package tree

import (
	"net/http"
	"testing"
)

func TestNewTree(t *testing.T) {
	tree := newTree()
	if tree.Index != 0 {
		t.Errorf("excpected 0, actual %v", tree.Index)
	}

	tree = newTree()
	if tree.Index != 0 {
		t.Errorf("expected 0, actual %v", tree.Index)
	}
}

func TestNewNode(t *testing.T) {
	tree := newTree()

	if tree.Index != 0 {
		t.Errorf("expected 0, actual %v", tree.Index)
	}

	node := tree.newNode(nil)

	if tree.Index != 1 {
		t.Errorf("expected 1, actual %v", tree.Index)
	}

	if node.Index != 1 {
		t.Errorf("expected 1, actual %v", node.Index)
	}

	node2 := node.NewNode()

	if node2.tree != tree {
		t.Errorf("expected %v, actual %v", tree, node2.tree)
	}

	if node2.Index != 2 {
		t.Errorf("expected 2, actual %v", node2.Index)
	}

	if node2.Parent != node {
		t.Errorf("expected %v, actual %v", node, node2.Parent)
	}

	if node2.Left != nil {
		t.Errorf("expected nil, actual %v", node2.Left)
	}

	if node2.Center != nil {
		t.Errorf("expected nil, actual %v", node2.Center)
	}

	if node2.Right != nil {
		t.Errorf("expected nil, actual %v", node2.Right)
	}
}

func TestDig(t *testing.T) {
	tree := newTree()
	node := tree.newNode(nil)

	result := node.Dig()

	if node.Left == nil {
		t.Errorf("not expected nil")
	}

	if node.Center == nil {
		t.Errorf("not expected nil")
	}

	if node.Right == nil {
		t.Errorf("node expected nil")
	}

	if result[left] != node.Left {
		t.Errorf("expected %v, actual %v", result[left], node.Left)
	}

	if result[center] != node.Center {
		t.Errorf("expected %v, actual %v", result[center], node.Center)
	}

	if result[right] != node.Right {
		t.Errorf("expected %v, actual %v", result[right], node.Right)
	}
}

func TestDepth(t *testing.T) {
	tree := newTree()
	node := tree.newNode(nil)

	if node.Depth() != 0 {
		t.Errorf("expected 0, actual %v", node.Depth())
	}

	node.Dig()

	if node.Left.Depth() != 1 {
		t.Errorf("expected 1, actual %v", node.Left.Depth())
	}

	node.Left.Dig()

	if node.Left.Left.Depth() != 2 {
		t.Errorf("expected 2, actual %v", node.Left.Left.Depth())
	}
}

func TestChildren(t *testing.T) {
	tree := newTree()
	node := tree.newNode(nil)

	node.Dig()

	children := node.Children()

	if children[left] != node.Left {
		t.Errorf("expected %v, actual %v", node.Left, children[left])
	}

	if children[center] != node.Center {
		t.Errorf("expected %v, actual %v", node.Center, children[center])
	}

	if children[right] != node.Right {
		t.Errorf("expected %v, actual %v", node.Right, children[right])
	}
}

func TestDig2(t *testing.T) {
	tree := newTree()
	node := tree.newNode(nil)

	result := Dig(node)

	if result[left].Index != 2 {
		t.Errorf("expected 2, actual %v", result[left].Index)
	}

	if result[center].Index != 3 {
		t.Errorf("expected 2, actual %v", result[center].Index)
	}

	if result[right].Index != 4 {
		t.Errorf("expected 2, actual %v", result[right].Index)
	}

	result = Dig(result...)

	if result[0].Index != 5 {
		t.Errorf("expected 5, actual %v", result[0].Index)
	}

	if result[4].Index != 9 {
		t.Errorf("expected 9, actual %v", result[4].Index)
	}

	if result[8].Index != 13 {
		t.Errorf("expected 13, actual %v", result[8].Index)
	}
}

func TestNew(t *testing.T) {
	node := New(3)

	if node.Left.Index != 2 {
		t.Errorf("expected 2, actual %v", node.Left.Index)
	}

	if node.Center.Index != 3 {
		t.Errorf("expected 2, actual %v", node.Center.Index)
	}

	if node.Right.Index != 4 {
		t.Errorf("expected 2, actual %v", node.Right.Index)
	}

	if node.Left.Left.Index != 5 {
		t.Errorf("expected 5, actual %v", node.Left.Left.Index)
	}

	if node.Center.Center.Index != 9 {
		t.Errorf("expected 9, actual %v", node.Center.Center.Index)
	}

	if node.Right.Right.Index != 13 {
		t.Errorf("expected 13, actual %v", node.Right.Right.Index)
	}

	if node.Right.Right.Right == nil {
		t.Errorf("expected nil")
	}

	if node.Left.Left.Left == nil {
		t.Errorf("expected nil")
	}

	if node.Center.Center.Center == nil {
		t.Errorf("expected nil")
	}
}

type responseWriter struct {
}

func (w responseWriter) Header() http.Header {
	return nil
}

func (w responseWriter) Write(b []byte) (int, error) {
	return 0, nil
}

func (w responseWriter) WriteHeader(i int) {

}

func TestServeHTTP(t *testing.T) {
	tree := newTree()
	node := tree.newNode(nil)

	node.ServeHTTP(responseWriter{}, nil)
}

func TestFind(t *testing.T) {
	node := New(3)

	if node.Find(2) != node.Left {
		t.Errorf("expected %v, actual %v", node.Left, node.Find(2))
	}

	if node.Find(3) != node.Center {
		t.Errorf("expected %v, actual %v", node.Center, node.Find(3))
	}

	if node.Find(4) != node.Right {
		t.Errorf("expected %v, actual %v", node.Right, node.Find(4))
	}

	if node.Find(5) != node.Left.Left {
		t.Errorf("expected %v, actual %v", node.Left.Left, node.Find(5))
	}

	if node.Find(9) != node.Center.Center {
		t.Errorf("expected %v, actual %v", node.Center.Center, node.Find(9))
	}

	if node.Find(13) != node.Right.Right {
		t.Errorf("expected %v, actual %v", node.Right.Right, node.Find(13))
	}

	if node.Find(44) != nil {
		t.Errorf("expected nil")
	}
}

func TestConFind(t *testing.T) {
	node := New(3)

	if node.ConFind(2) != node.Left {
		t.Errorf("expected %v, actual %v", node.Left.Index, node.ConFind(2))
	}

	if node.ConFind(3) != node.Center {
		t.Errorf("expected %v, actual %v", node.Center.Index, node.ConFind(3))
	}

	if node.ConFind(4) != node.Right {
		t.Errorf("expected %v, actual %v", node.Right.Index, node.ConFind(4))
	}

	if node.ConFind(5) != node.Left.Left {
		t.Errorf("expected %v, actual %v", node.Left.Left.Index, node.ConFind(5))
	}

	if node.ConFind(9) != node.Center.Center {
		t.Errorf("expected %v, actual %v", node.Center.Center.Index, node.ConFind(9))
	}

	if node.ConFind(13) != node.Right.Right {
		t.Errorf("expected %v, actual %v", node.Right.Right.Index, node.ConFind(13))
	}

	if node.ConFind(44) != nil {
		t.Errorf("expected nil")
	}
}

func BenchmarkFind(b *testing.B) {
	node := New(3)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		node.Find(2)
		node.Find(3)
		node.Find(4)
		node.Find(5)
		node.Find(9)
		node.Find(13)
		node.Find(44)
	}
}

func BenchmarkConFind(b *testing.B) {
	node := New(3)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		node.ConFind(2)
		node.ConFind(3)
		node.ConFind(4)
		node.ConFind(5)
		node.ConFind(9)
		node.ConFind(13)
		node.ConFind(44)
	}
}
