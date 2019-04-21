package main

import (
	"awesomeProject/tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

// 后序遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Left.Left = tree.CreateTreeNodes(7)
	root.Left.Right = tree.CreateTreeNodes(10)
	root.Right = &tree.Node{Value: 5, Left: nil, Right: nil}
	root.Right.Left = new(tree.Node)
	root.Right.Right = tree.CreateTreeNodes(4)
	root.Traverse()
	fmt.Println()

	// 扩展postOrder方法
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	// 函数式编程 遍历二叉树
	fmt.Printf("root has %d node" , root.NodeCount())
}
