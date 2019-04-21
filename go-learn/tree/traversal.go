package tree

import "fmt"

func (node *Node) Traverse() {
	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
	fmt.Println()
}

func (node *Node) NodeCount() int  {
	counter := 0
	node.TraverseFunc(func(node *Node) {
		counter++
	})
	return counter
}

func (node *Node) TraverseFunc(f func(*Node))  {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}


