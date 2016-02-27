package tree

import (
	"strconv"
	"strings"
)

func read(str string) (int, int) {
	arr := strings.Split(str, "->")
	s, err := strconv.Atoi(arr[0])
	if err != nil {
		panic(err)
	}
	e, err := strconv.Atoi(arr[1])
	if err != nil {
		panic(err)
	}
	return s, e
}

func isMe(a *Node, b *Node) bool {
	return a == b
}

func isMother(a *Node, b *Node) bool {
	return a.Parent == b
}

func isAunt(a *Node, b *Node) bool {
	c := a.Parent
	if c == nil {
		return false
	}
	return c.Parent == b.Parent
}

func isDaughter(a *Node, b *Node) bool {
	return a == b.Parent
}

func isSister(a *Node, b *Node) bool {
	return a.Parent == b.Parent
}

func isNiece(a *Node, b *Node) bool {
	c := b.Parent
	if c == nil {
		return false
	}
	return a.Parent == c.Parent
}

func isCo(a *Node, b *Node) bool {
	c := a.Parent
	d := b.Parent
	if c == nil || d == nil {
		return false
	}
	return c.Parent == d.Parent
}

func solve(str string) string {
	node := New(3)
	s, e := read(str)
	sNode := node.Find(s)
	eNode := node.Find(e)

	if isMe(sNode, eNode) {
		return "me"
	}
	if isMother(sNode, eNode) {
		return "mo"
	}
	if isAunt(sNode, eNode) {
		return "au"
	}
	if isDaughter(sNode, eNode) {
		return "da"
	}
	if isSister(sNode, eNode) {
		return "si"
	}
	if isNiece(sNode, eNode) {
		return "ni"
	}
	if isCo(sNode, eNode) {
		return "co"
	}

	return "-"
}
