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

func isDaughter(a *Node, b *Node) bool {
	return a == b.Parent
}

func isSister(a *Node, b *Node) bool {
	if isMe(a, b) {
		return false
	}
	return a.Parent == b.Parent
}

func isAunt(a *Node, b *Node) bool {
	if isMother(a, b) {
		return false
	}
	c := a.Parent
	if c == nil {
		return false
	}
	return c.Parent == b.Parent
}

func isNiece(a *Node, b *Node) bool {
	if isDaughter(a, b) {
		return false
	}
	c := b.Parent
	if c == nil {
		return false
	}
	return a.Parent == c.Parent
}

func isCo(a *Node, b *Node) bool {
	if isMe(a, b) {
		return false
	}
	if isSister(a, b) {
		return false
	}
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

func conSolve(str string) string {
	node := New(3)
	s, e := read(str)
	sNode := node.Find(s)
	eNode := node.Find(e)

	var result string
	resultCh := make(chan string, 1)
	failCh := make(chan bool, 7)
	failCount := 0

	go func(a *Node, b *Node) {
		if isMe(a, b) {
			resultCh <- "me"
		} else {
			failCh <- true
		}
	}(sNode, eNode)

	go func(a *Node, b *Node) {
		if isMother(a, b) {
			resultCh <- "mo"
		} else {
			failCh <- true
		}
	}(sNode, eNode)

	go func(a *Node, b *Node) {
		if isAunt(a, b) {
			resultCh <- "au"
		} else {
			failCh <- true
		}
	}(sNode, eNode)

	go func(a *Node, b *Node) {
		if isDaughter(a, b) {
			resultCh <- "da"
		} else {
			failCh <- true
		}
	}(sNode, eNode)

	go func(a *Node, b *Node) {
		if isSister(a, b) {
			resultCh <- "si"
		} else {
			failCh <- true
		}
	}(sNode, eNode)

	go func(a *Node, b *Node) {
		if isNiece(a, b) {
			resultCh <- "ni"
		} else {
			failCh <- true
		}
	}(sNode, eNode)

	go func(a *Node, b *Node) {
		if isCo(a, b) {
			resultCh <- "co"
		} else {
			failCh <- true
		}
	}(sNode, eNode)

	for {
		select {
		case result = <-resultCh:
			return result
		case <-failCh:
			failCount = failCount + 1
			if failCount > 6 {
				return "-"
			}
		}
	}
}
